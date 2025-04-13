package internal

import (
    "fmt"
    "os"
    "os/exec"
)

const (
    wordlistsDir = "./wordlists/"
    outputDir    = "./output/"
    domainListURL = "https://gist.github.com/jareddarkweb/64622715781e1637783e5c79823a3d99"
)

func RunWorkflow() error {
    os.MkdirAll(outputDir, 0755)

    domains, err := FetchDomains(domainListURL)
    if err != nil {
        return err
    }

    var finalOutputs []string

    for _, domain := range domains {
        fmt.Println("🔍 Enumerating:", domain)
        sub := runSubfinder(domain)
        goal := runGoaltdns(sub)
        got := runGotator(sub)
        pure := runPuredns(got)
        finalOutputs = append(finalOutputs, pure)
    }

    return mergeFiles(finalOutputs, outputDir+"result.txt")
}

func runCommand(name string, args ...string) string {
    fmt.Println("⚙️", name, args)
    cmd := exec.Command(name, args...)
    output := outputDir + name + ".txt"
    file, _ := os.Create(output)
    cmd.Stdout = file
    cmd.Stderr = os.Stderr
    _ = cmd.Run()
    return output
}

func runSubfinder(domain string) string {
    return runCommand("subfinder", "-d", domain, "-o", outputDir+domain+"_subs.txt")
}

func runGoaltdns(inputFile string) string {
    return runCommand("goaltdns", "-w", wordlistsDir+"goaltdns.txt", "-i", inputFile, "-o", inputFile+".goaltdns.txt")
}

func runGotator(inputFile string) string {
    return runCommand("gotator", "-sub", inputFile, "-perm", wordlistsDir+"goaltdns.txt", "-depth", "2", "-adv", "-silent", "-o", inputFile+".gotator.txt")
}

func runPuredns(inputFile string) string {
    return runCommand("puredns", "resolve", inputFile, "-r", "resolvers.txt", "-o", inputFile+".resolved.txt")
}
