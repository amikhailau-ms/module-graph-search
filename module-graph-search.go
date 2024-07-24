package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var (
	dependenciesToSearchFor []string
	dependenciesRaw         string
)

type dependencyGraph struct {
	verticesMap    map[string]int
	vertices       []string
	allEdges       [][]int
	edgesByVertice map[int][]int
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: go mod graph | module-graph-search --deps=module1[,module2,...]`)
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("module-graph-searcher: ")

	flag.StringVar(&dependenciesRaw, "deps", "", "comma separated list of dependencies to search for")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
	}

	if dependenciesRaw != "" {
		dsf := strings.Split(dependenciesRaw, ",")
		for _, dependency := range dsf {
			ds := strings.TrimSpace(dependency)
			if ds != "" {
				dependenciesToSearchFor = append(dependenciesToSearchFor, ds)
			}
		}
	}

	if err := searchForDependency(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func searchForDependency(in io.Reader, out io.Writer) error {
	dg, err := buildReverseDependencyGraph(in)
	if err != nil {
		return err
	}

	for i, dependencyToSearch := range dependenciesToSearchFor {
		if i != 0 {
			fmt.Fprintln(out)
			fmt.Fprintln(out)
		}
		result := dg.findFullDependencyChains(dependencyToSearch)
		fmt.Fprintf(out, Cyan+"Dependency chains for %s\n\n", dependencyToSearch)
		for _, chain := range result {
			for i, module := range chain {
				if i != 0 {
					fmt.Fprintf(out, "<--")
				} else {
					fmt.Fprintf(out, Purple)
				}
				fmt.Fprintf(out, "%s", module)
			}
			fmt.Fprintln(out)
		}
	}

	return nil
}

func buildReverseDependencyGraph(in io.Reader) (*dependencyGraph, error) {
	dg := dependencyGraph{
		verticesMap:    make(map[string]int),
		edgesByVertice: make(map[int][]int),
	}
	scanner := bufio.NewScanner(in)
	currSize := 0

	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		parts := strings.Fields(l)
		if len(parts) != 2 {
			return nil, fmt.Errorf("failed to parse graph edge: %s", l)
		}
		from := parts[0]
		to := parts[1]
		fromIndex, ok := dg.verticesMap[from]
		if !ok {
			dg.verticesMap[from] = currSize
			fromIndex = currSize
			dg.vertices = append(dg.vertices, from)
			currSize += 1
		}
		toIndex, ok := dg.verticesMap[to]
		if !ok {
			dg.verticesMap[to] = currSize
			toIndex = currSize
			dg.vertices = append(dg.vertices, to)
			currSize += 1
		}

		dg.allEdges = append(dg.allEdges, []int{toIndex, fromIndex})
		dg.edgesByVertice[toIndex] = append(dg.edgesByVertice[toIndex], fromIndex)
	}
	return &dg, nil
}

func (dg *dependencyGraph) findFullDependencyChains(dependency string) [][]string {
	dIndex, ok := dg.verticesMap[dependency]
	if !ok {
		return nil
	}
	var result [][]string

	queue := []int{dIndex}
	visited := make(map[int][]int)
	visited[dIndex] = []int{dIndex}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nextVertices := dg.edgesByVertice[node]
		currPath := visited[node]
		for _, nv := range nextVertices {
			if _, ok := visited[nv]; ok {
				continue
			}
			nvName := dg.vertices[nv]
			if !strings.Contains(nvName, "@") {
				// reached root node - means we have found the full dependency chain
				newChain := make([]string, 0, len(currPath))
				for _, pathIndex := range currPath {
					newChain = append(newChain, dg.vertices[pathIndex])
				}
				result = append(result, newChain)
				continue
			}
			pathCopy := make([]int, len(currPath), len(currPath)+1)
			copy(pathCopy, currPath)
			pathCopy = append(pathCopy, nv)
			visited[nv] = pathCopy
			queue = append(queue, nv)
		}
	}

	return result
}
