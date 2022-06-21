package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	//"image/color"
	graph "kubernetes-go/graph"
	"math"
	//"os/exec"

	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/widget"
	"github.com/aquasecurity/table"
)

// func makeTable(data [][]string) *widget.Table{
// 	newlist := widget.NewTable(
// 		func() (int, int) {
// 			return len(data), len(data[0])
// 		},
// 		func() fyne.CanvasObject {
// 			return widget.NewLabel("wide content")
// 		},
// 		func(i widget.TableCellID, o fyne.CanvasObject) {
// 			o.(*widget.Label).SetText(data[i.Row][i.Col])
// 		})
// 	return newlist
// }
/*
func main(){
	myApp := app.New()
	myWindow := myApp.NewWindow("Kubernetes metrics")
	myWindow.Resize(fyne.NewSize(850, 950))
	title := canvas.NewText("Get graphs for pods and nodes", color.White)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24
	namespace := widget.NewEntry()
	namespace.SetPlaceHolder("Enter namespace...")
	entity := widget.NewEntry()
	entity.SetPlaceHolder("Enter entity name... either pods or nodes")

	// if fmt.Sprintf("%v",entity) == fmt.Sprintf("%v", "pods"){
	// 	graph.GenerateGraphPods(fmt.Sprintf("%v", namespace))
	// }
	logo := canvas.NewImageFromFile("1200px-Fyne_toolkit_logo.svg.png")
	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord
	btn := widget.NewButton("Show graph", nil)
	changed := func(){
		fmt.Println(namespace.Text)
		fmt.Println(entity.Text)
		if entity.Text == "pods"{
			data := graph.GenerateGraphPods(namespace.Text)
			//list  = makeTable(data)
			words := ""
			words = words + "Pod Name" +"\t\t\t\t\t" +"Container Name" + "\t\t\t\t\t"+"CPU Percentage" +"\t\t\t\t\t"+ "Memory Percentage"+"\n\n"
			for _, c := range data{
				words = words + c[0]+"\t\t"+c[1]+"\t\t"+c[2]+"\t\t"+c[3] + "\n\n"

			}
			factText.SetText(words)
			// img := canvas.NewImageFromFile("cpu_consumerclf.png")
			// img1 := canvas.NewImageFromFile("memory_consumerclf.png")
			// imagesGraph = container.NewVSplit(img, img1)

		}
	}
	btn.OnTapped = changed



	vBox := container.New(layout.NewVBoxLayout(), title, namespace,entity, btn, factText)
	myWindow.SetContent(container.NewHSplit(vBox, logo))
	myWindow.ShowAndRun()
	cmd := exec.Command("feh", "output1.png")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout))

}
*/

func main() {
	
	namespace := flag.String("namespace", "", "Enter Namespace")
	entity := flag.String("entity", "", "Enter pods/ nodes (For nodes only use entity flag) (Required)")
	help := flag.String("h", "","shows Usage of the command line arguments")

	flag.Parse()
	//if help == ""
	// namespace := os.Args[1]
	// entity := os.Args[2]
	if *help == ""{
		flag.PrintDefaults()
	}
	if *namespace == "" {
		fmt.Println("Namespace not needed for nodes")
	}
	if *entity == "pods" {
		start := time.Now()
		fmt.Println("Generating metrics.............")
		data, _ := graph.GetPodMetrics(*namespace)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Println("Time taken to get the metrics: ", math.Round(elapsed.Seconds()), "s")
		t := table.New(os.Stdout)

		t.SetHeaders("Node Name", "Pod Name", "Container Name", "CPU value", "CPU Percentage", "Memory Value", "Memory Percentage")
		for _, v := range data {
			t.AddRow(v[0], v[1], v[2], v[3], v[4], v[5], v[6])
		}
		t.Render()

	} else if *entity == "nodes" {
		start := time.Now()
		fmt.Println("Generating metrics.............")
		data := graph.GetNodeMetrics()
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Println("Time taken to get the metrics: ", math.Round(elapsed.Seconds()), "s")
		t := table.New(os.Stdout)

		t.SetHeaders("Node Name", "CPU Value", "Node CPU usage", "CPU Percentage", "Memory Value", "Node Memory usage", "Memory Percentage", "No of pods")
		for _, v := range data {
			t.AddRow(v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])
		}
		t.Render()

	}

}
