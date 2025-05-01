package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	namespace   string
	pod         string
	follow      bool
	tail        int64
	all         bool
	search      string
	searchTerms []string
)

// Color definitions
var (
	headerColor    = color.New(color.FgHiCyan, color.Bold)
	debugColor     = color.New(color.FgHiBlue)
	infoColor      = color.New(color.FgHiGreen)
	warnColor      = color.New(color.FgHiYellow)
	errorColor     = color.New(color.FgHiRed)
	timestampColor = color.New(color.FgHiWhite)
	searchColors   = []*color.Color{
		color.New(color.BgYellow, color.FgBlack),
		color.New(color.BgGreen, color.FgBlack),
		color.New(color.BgBlue, color.FgBlack),
		color.New(color.BgMagenta, color.FgBlack),
	}
)

func printSectionSeparator() {
	fmt.Println("\n" + strings.Repeat("═", 100) + "\n")
}

func printHeader(text string) {
	headerColor.Println("\n" + strings.Repeat("═", 100))
	headerColor.Printf("║ %s\n", text)
	headerColor.Println(strings.Repeat("═", 100))
}

func highlightSearchTerms(message string) string {
	for i, term := range searchTerms {
		if i < len(searchColors) {
			message = strings.ReplaceAll(message, term, searchColors[i].Sprintf(term))
		}
	}
	return message
}

func printLogEntry(timestamp, level, message string) {
	// Format timestamp
	t, err := time.Parse(time.RFC3339, timestamp)
	if err == nil {
		timestamp = t.Format("2006-01-02 15:04:05.000")
	}

	// Try to parse JSON message
	var jsonData interface{}
	if err := json.Unmarshal([]byte(message), &jsonData); err == nil {
		// It's JSON, pretty print it
		prettyJSON, _ := json.MarshalIndent(jsonData, "  ", "  ")
		message = string(prettyJSON)
	}

	// Highlight search terms if present
	if len(searchTerms) > 0 {
		message = highlightSearchTerms(message)
	}

	// Choose color based on log level
	var levelColor *color.Color
	switch strings.ToUpper(level) {
	case "DEBUG":
		levelColor = debugColor
	case "INFO":
		levelColor = infoColor
	case "WARN":
		levelColor = warnColor
	case "ERROR":
		levelColor = errorColor
	default:
		levelColor = color.New(color.FgWhite)
	}

	// Print log entry with proper formatting
	fmt.Printf("%s │ %s │ %s\n",
		timestampColor.Sprintf("%-23s", timestamp),
		levelColor.Sprintf("%-7s", level),
		message)
}

func shouldShowLog(line string) bool {
	if len(searchTerms) > 0 {
		for _, term := range searchTerms {
			if strings.Contains(line, term) {
				return true
			}
		}
		return false
	}
	if !all {
		return strings.Contains(strings.ToLower(line), "debug")
	}
	return true
}

func getLogLevel(message string) string {
	messageLower := strings.ToLower(message)
	if strings.Contains(messageLower, "error") {
		return "ERROR"
	} else if strings.Contains(messageLower, "warn") {
		return "WARN"
	} else if strings.Contains(messageLower, "debug") {
		return "DEBUG"
	}
	return "INFO"
}

var rootCmd = &cobra.Command{
	Use:   "podlog",
	Short: "A modern CLI tool for viewing Kubernetes pod debug logs",
	Long: `podlog is a CLI tool that makes it easy to view and analyze debug logs from Kubernetes pods.
It provides beautiful formatting for JSON structures and focuses on debug statements.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse search terms
		if search != "" {
			searchTerms = strings.Split(search, ",")
			for i := range searchTerms {
				searchTerms[i] = strings.TrimSpace(searchTerms[i])
			}
		}

		// Get Kubernetes client
		clientset, err := getClientSet()
		if err != nil {
			return fmt.Errorf("error creating Kubernetes client: %w", err)
		}

		// Get the pod
		p, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), pod, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("error getting pod %s in namespace %s: %w", pod, namespace, err)
		}

		// Print pod info header
		printHeader(fmt.Sprintf("📦 Pod: %s", p.Name))
		fmt.Printf("Container: %s\n", p.Spec.Containers[0].Name)
		if len(searchTerms) > 0 {
			fmt.Println("Searching for:")
			for i, term := range searchTerms {
				if i < len(searchColors) {
					fmt.Printf("  %s\n", searchColors[i].Sprintf(term))
				}
			}
		}
		printSectionSeparator()

		opts := &corev1.PodLogOptions{
			Follow:    follow,
			TailLines: &tail,
			Container: p.Spec.Containers[0].Name,
		}

		req := clientset.CoreV1().Pods(namespace).GetLogs(pod, opts)
		stream, err := req.Stream(context.Background())
		if err != nil {
			return fmt.Errorf("error getting logs for pod %s: %w", pod, err)
		}
		defer stream.Close()

		// Process log stream
		scanner := bufio.NewScanner(stream)
		for scanner.Scan() {
			line := scanner.Text()

			if !shouldShowLog(line) {
				continue
			}

			// Try to parse the log line
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				timestamp := parts[0]
				message := parts[1]
				level := getLogLevel(message)
				printLogEntry(timestamp, level, message)
			} else {
				printLogEntry("", "LOG", line)
			}
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading log stream: %w", err)
		}

		return nil
	},
}

// getClientSet initializes a Kubernetes clientset (supports in-cluster and local config)
func getClientSet() (*kubernetes.Clientset, error) {
	// Try in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to kubeconfig for local dev
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace")
	rootCmd.Flags().StringVarP(&pod, "pod", "p", "", "Pod name")
	rootCmd.Flags().BoolVarP(&follow, "follow", "f", false, "Follow log output")
	rootCmd.Flags().Int64VarP(&tail, "tail", "t", 100, "Number of lines to show from the end of the logs")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "Show all logs (not just debug)")
	rootCmd.Flags().StringVarP(&search, "search", "s", "", "Search for specific text in logs (comma-separated for multiple terms)")

	rootCmd.MarkFlagRequired("namespace")
	rootCmd.MarkFlagRequired("pod")
}
