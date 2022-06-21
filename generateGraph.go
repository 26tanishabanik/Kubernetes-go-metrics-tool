package graph

import (
	"context"
	// "encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	//"math/rand"
	"github.com/wcharczuk/go-chart/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// corev1 "k8s.io/api/core/v1"
	// histogram "kubernetes-go/histogram"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type PodMetricsList struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink string `json:"selfLink"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			SelfLink          string    `json:"selfLink"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Timestamp  time.Time `json:"timestamp"`
		Window     string    `json:"window"`
		Containers []struct {
			Name  string `json:"name"`
			Usage struct {
				CPU    string `json:"cpu"`
				Memory string `json:"memory"`
			} `json:"usage"`
		} `json:"containers"`
	} `json:"items"`
}

// func getMetrics(clientset *kubernetes.Clientset, pods *PodMetricsList) error {
// 	data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
// 	if err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal(data, &pods)
// 	return err
// }

/*
func main() {
	namespace := os.Args[1]
	fmt.Println(namespace)
	// var kubeconfig, master string
	// config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
    // if err != nil{
    //     panic(err)
    // }

    // mc, err := metrics.NewForConfig(config)
    // if err != nil {
    //     panic(err)
    // }
	// fmt.Println(mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(context.Background(),metav1.ListOptions{}))
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Printf("Error in new client config: %s\n",err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	// *********** client set for metrics *********
	clientsetMetrics, err := metrics.NewForConfig(config)
	if err != nil {
		log.Println("Error in new config: \n",err)
	}
	// ************** node metrics *************
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting pods list for namespace %s: %s\n",namespace, err)
	}
	for i, n := range nodeList.Items {
		fmt.Printf("Pod %d in %s namespace: %s \n",i, namespace,n.Name)
	}
	nodeMetricsList, err := clientsetMetrics.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error in getting pod metrics for namespace %s: %s\n",namespace, err)
	}
	for _, v := range nodeMetricsList.Items {
		fmt.Printf("%s\n", v.GetName())
		fmt.Printf("%s\n", v.GetNamespace())
		fmt.Printf("%vm\n", v.Usage.Cpu().MilliValue())
		fmt.Printf("%vMi\n", v.Usage.Memory().Value()/(1024*1024))
	}
	// ************* pod metrics ************
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting pods list for namespace %s: %s\n",namespace, err)
	}
	for i, n := range podList.Items {
		fmt.Printf("Pod %d in %s namespace: %s \n",i, namespace,n.Name)
	}

	podMetricsList, err := clientsetMetrics.MetricsV1beta1().PodMetricses(namespace).List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Printf("Error in getting pod metrics for namespace %s: %s\n",namespace, err)
	}
	for _, v := range podMetricsList.Items {
		fmt.Printf("%s\n", v.GetName())
		fmt.Printf("%s\n", v.GetNamespace())
		fmt.Printf("%dm\n", v.Containers[0].Usage.Cpu().MilliValue())
		fmt.Printf("%dMi\n", v.Containers[0].Usage.Memory().Value()/(1024*1024))
	}

	// ********** new pod creation **********
	// newPod := &corev1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 	  Name: "test-pod",
	// 	},
	// 	Spec: corev1.PodSpec{
	// 	  Containers: []corev1.Container{
	// 		{Name: "busybox", Image: "busybox:latest",Command: []string{"sleep", "100000"}},
	// 	  },
	// 	},
	//   }

	// pod, err := clientset.CoreV1().Pods(namespace).Create(context.Background(), newPod, metav1.CreateOptions{})
	// if err != nil {
	// 	log.Printf("Error in creating pod for namespace %s: %s",namespace, err)
	// }
	// fmt.Println(pod)


	// var pods PodMetricsList
	// err = getMetrics(clientset, &pods)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// for _, m := range pods.Items {
	// 	fmt.Println(m.Metadata.Name, m.Metadata.Namespace, m.Timestamp.String())
	// }
	// mc, err := metrics.NewForConfig(configMetrics)
	// if err != nil {
	// 	log.Println(err)
	// }

	// podMetrics, err := mc.MetricsV1beta1().PodMetricses(namespace).List(metav1.ListOptions{})
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// for _, podMetric := range podMetrics.Items {
	// 	podContainers := podMetric.Containers
	// 	for _, container := range podContainers {
	// 		cpuQuantity, ok := container.Usage.Cpu().AsInt64()
	// 		memQuantity, ok := container.Usage.Memory().AsInt64()
	// 		if !ok {
	// 			return
	// 		}
	// 		msg := fmt.Sprintf("Container Name: %s \n CPU usage: %d \n Memory usage: %d", container.Name, cpuQuantity, memQuantity)
	// 		fmt.Println(msg)
	// 	}

	// }

}

*/
/*
func main(){
	//namespace := "cainaticsmongo"
    // nodeName := "aks-memorypool-30666732-vmss000002"
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Printf("Error in new client config: %s\n",err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
    // events, err := clientset.CoreV1().Events(namespace).List(context.Background(), metav1.ListOptions{
    //     TypeMeta: metav1.TypeMeta{
    //         Kind: "Node",
    //     },
    //     FieldSelector: "involvedObject.name=" + nodeName,
    // })
	// fmt.Println(events)
	podList, err := clientset.CoreV1().Pods("jenkins").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting pods list for namespace %s: %s\n","jenkins", err)
	}
	fmt.Println(len(podList.Items))
	// for _, n := range podList.Items {
	// 	fmt.Println(n.Spec.NodeName)
	// }
	// nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

    // if err != nil {
    //     fmt.Println(err)
    // }
	// for _, v := range nodes.Items{
	// 	if v.Name == "aks-memorypool-30666732-vmss000001"{

	// 		fmt.Println(v.Status.Capacity.Cpu())
	// 	}
	// }

}
*/
type podGraph struct {
	nodeName string
	containerName string
	podName string
	containerCpu float64
	cpuPercentage  float64
	containerMemory float64
	memoryPercentage  float64

}

type MyBox struct {
	Graph []podGraph
}
type nodeGraph struct {
	nodeName string
	nodeCpuPercentage  float64
	nodeMemory float64
	nodeMemoryPercentage  float64
	noOfPods int
}

type MyNode struct {
	Graph []nodeGraph
}

func clientSetup() *kubernetes.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Printf("Error in new client config: %s\n", err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset

}
func clientMetricsSetup() *metrics.Clientset {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Printf("Error in new client config: %s\n", err)
	}
	clientsetMetrics, err := metrics.NewForConfig(config)
	if err != nil {
		log.Printf("Error in new client metrics config: %s\n", err)
	}
	return clientsetMetrics

}

func GetNodeCapacity(nodeName string) (int64, *resource.Quantity) {
	clientset := clientSetup()
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Println(err)
	}
	for _, v := range nodes.Items {
		
		if v.Name == nodeName {
			// fmt.Println(v.Name)
			// fmt.Println("Node CPU Value: ", v.Status.Capacity.Cpu())
			// fmt.Println("Node CPU MiliValue: ", v.Status.Capacity.Cpu().MilliValue())
			// fmt.Println("Node Memory: ", v.Status.Capacity.Memory())
			// fmt.Println("Node Memory Value: ", v.Status.Capacity.Memory().Value())
			return v.Status.Capacity.Cpu().MilliValue(), v.Status.Capacity.Memory()
		}
	}
	return -1, nil

}
func GetNodeName(podName string, namespace string) (nodeName string){
	clientset := clientSetup()
	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting pods list for namespace %s: %s\n", namespace, err)
	}
	for _, n := range podList.Items {
		if n.Name == podName{
			return n.Spec.NodeName
		}
	}
	return ""
}
func GetNodeMetrics() ([][]string){
	var data [][]string
	clientset := clientSetup()
	clientsetMetrics := clientMetricsSetup()
	nodeMetricsList, err := clientsetMetrics.MetricsV1beta1().NodeMetricses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting nodes list : %s\n", err)
	}
	// nodeBox := MyNode{}
	for _, n := range nodeMetricsList.Items {
		nodeName := n.Name
		
		cpu, memory := GetNodeCapacity(nodeName)
		cpuFloat := float64(cpu * 1000.0)
		memoryValue:= n.Usage.Memory()
		memoryFloat := float64(memory.Value()/(1024 * 1024))
		nodeCpuUsage := float64(n.Usage.Cpu().MilliValue()) * 1000
		nodeMemoryUsage := float64(n.Usage.Memory().Value())/(1024.0 * 1024)
		// numberPods := n.Usage.Pods().Size()
		pods, _:= clientset.CoreV1().Pods("").List(context.Background(),metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + nodeName,
		})
		numberPods := len(pods.Items)
		// item := nodeGraph{nodeName: nodeName, nodeCpuPercentage: (nodeCpuUsage/(cpuFloat)) * 100, nodeMemory: memoryFloat,nodeMemoryPercentage: nodeMemoryUsage/(memoryFloat) * 100, noOfPods: numberPods}
		// nodeBox.Graph = append(nodeBox.Graph, item)
		data = append(data, []string{nodeName, fmt.Sprintf("%v", cpu),fmt.Sprintf("%v",n.Usage.Cpu().MilliValue()) ,fmt.Sprintf("%v",(nodeCpuUsage/(cpuFloat)) * 100), fmt.Sprintf("%v", memory),fmt.Sprintf("%v",memoryValue),fmt.Sprintf("%v", nodeMemoryUsage/(memoryFloat) * 100), fmt.Sprintf("%v",numberPods)})

	}
	return data

}
func GetPodMetrics(namespace string) ([][]string, MyBox){
	clientsetMetrics := clientMetricsSetup()
	podMetricsList, err := clientsetMetrics.MetricsV1beta1().PodMetricses(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error in getting pods list for namespace %s: %s\n", namespace, err)
	}
	box := MyBox{}
	var data [][]string
	for _, n := range podMetricsList.Items {
		nodeName := GetNodeName(n.Name, namespace)
		cpu, memory := GetNodeCapacity(nodeName)
		cpuFloat := float64(cpu)
		memoryFloat := float64(memory.Value()/(1024 * 1024))
		for _, container := range n.Containers{	
			containerCpuFloat:= float64(container.Usage.Cpu().MilliValue())
			containerMemoryFloat := float64(container.Usage.Memory().Value())/(1024.0 * 1024.0)
			// item := podGraph{nodeName: nodeName, containerName: container.Name, podName: n.Name, containerCpu: containerCpuFloat, cpuPercentage: (containerCpuFloat/(cpuFloat)) * 100, containerMemory: containerMemoryFloat, memoryPercentage: containerMemoryFloat/(memoryFloat) * 100}
			// box.Graph = append(box.Graph, item)
			data = append(data, []string{nodeName, n.Name, container.Name, fmt.Sprintf("%v",containerCpuFloat), fmt.Sprintf("%v", (containerCpuFloat/cpuFloat) * 100), fmt.Sprintf("%v", containerMemoryFloat), fmt.Sprintf("%v",(containerMemoryFloat/memoryFloat) * 100)})
		}
	}
	//fmt.Println(box)
	return data, box

}
func GenerateGraphPods(namespace string) ([][]string){
	start := time.Now()
	fmt.Println("Generating metrics.............")
	_, graphValues := GetPodMetrics(namespace)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Time taken to get the metrics: ",elapsed)
	var cpuPie []chart.Value
	var memoryPie []chart.Value
	var data [][]string
	//colors := [...]string{"Blue", "Green", "Orange", "Gray", "Yellow", "Cyan", "White", "Black", "LightGray", "AlternateLightGray", "AlternateGreen", "AlternateBlue", "AlternateYellow", "AlternateGray"}
	for _, v := range graphValues.Graph{
		podContainerName := v.podName + "\n" +v.containerName
		cpuPie = append(cpuPie, chart.Value{Value:v.cpuPercentage,Label: podContainerName + fmt.Sprintf("%v", v.cpuPercentage)})
		memoryPie = append(memoryPie, chart.Value{Value:v.memoryPercentage,Label: podContainerName + fmt.Sprintf("%v", v.memoryPercentage)})
		data = append(data, []string{v.nodeName, v.podName, v.containerName, fmt.Sprintf("%v",v.cpuPercentage), fmt.Sprintf("%v", v.memoryPercentage)})
	}
	cpuPieChart := chart.PieChart{
		Width:  900,
		Height: 1024,
		Values: cpuPie,
		DPI: 40,
		Canvas: chart.Style{FontSize: 10},
	}
	memoryPieChart := chart.PieChart{
		Width:  900,
		Height: 1024,
		Values: memoryPie,
		DPI: 40,
		Canvas: chart.Style{FontSize: 10},
	}
	


	cpu, _ := os.Create("cpu_"+namespace+".png")
	memory, _ := os.Create("memory_"+namespace+".png")
	fmt.Println("Generated metrics")
	defer cpu.Close()
	defer memory.Close()
	cpuPieChart.Render(chart.PNG, cpu)
	memoryPieChart.Render(chart.PNG, memory)
	return data

}

func barPlot(values plotter.Values) {
    p:= plot.New()
    
    p.Title.Text = "bar plot"
 
    bar, err := plotter.NewBarChart(values, 15)
    if err != nil {
        panic(err)
    }
    p.Add(bar)
 
    if err := p.Save(3*vg.Inch, 3*vg.Inch, "bar.png"); err != nil {
        panic(err)
    }
}
func histPlot(values plotter.Values) {
    p:= plot.New()
    p.Title.Text = "histogram plot"
 
    hist, err := plotter.NewHist(values, 20)
    if err != nil {
        panic(err)
    }
    p.Add(hist)
 
    if err := p.Save(3*vg.Inch, 3*vg.Inch, "hist.png"); err != nil {
        panic(err)
    }
}


