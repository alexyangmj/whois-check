package main

import (
	"fmt"
    "os"
    "net"
    "bufio"
    "strings"
	"github.com/likexian/whois"
	"github.com/likexian/whois-parser"
    
)

func IsIpv4Net(host string) bool {
   return net.ParseIP(host) != nil
}

func IsIpv6Net(host string) bool {
   return net.ParseIP(host) != nil
}

func main() {
    
    Banner := "whois-check v2.0\n"
    Banner = Banner + "Last Update: 14 Apr 2024, Alex Yang (https://linkedin.com/in/4yang)\n\n"
    Banner = Banner + "Usage for Single IP query:\n"
    Banner = Banner + "    whois-check [ipv4 | ipv6 | domain.com]\n\n"
    Banner = Banner + "Optional_Switch for output format (FOR DOMAIN ONLY):\n"
    Banner = Banner + "    S   Show only resolved IP & domain status\n"    
    Banner = Banner + "    v   Show verbose (more) information\n"
    Banner = Banner + "    N   Show only resolved IP, registrant name, organization & country\n"
    Banner = Banner + "    T   Show only created, updated and expiration date\n\n"
    Banner = Banner + "Example:\n"
    Banner = Banner + "   whois-check twitter.com\n"
    Banner = Banner + "   whois-check google.com c\n"
    Banner = Banner + "   whois-check youtube.com N\n"
    Banner = Banner + "   whois-check netflix.com T\n\n"
    Banner = Banner + "Optional_Switch for output format (FOR IPv4/v6 ONLY):\n"
    Banner = Banner + "    C   Show only CIDR\n"    
    Banner = Banner + "    R   Show only Cannonical Name\n\n"
    Banner = Banner + "Example:\n"
    Banner = Banner + "   whois-check 20.231.239.246 C\n"
    Banner = Banner + "   whois-check 142.251.12.94 R\n\n"
    Banner = Banner + "Usage for Bulk IP query:\n"
    Banner = Banner + "   whois-check inputfile.txt\n"
    Banner = Banner + "   (the input filename must be STRICTLY inputfile.txt. Lines must only a IPv4/6 and domain name)\n"

    var input       string
    var isFile      bool   = false
    var Switch      string = "NIL"
    
    errline := 1
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Can't process further, likely a bug")
        }
    }()

    if len(os.Args) == 1  { 
        fmt.Println(Banner)
        return
    } 
    
    if len(os.Args) == 2 { input = os.Args[1] }
    
    if len(os.Args) == 3 { 
        input   = os.Args[1]
        Switch  = os.Args[2] 
    }
    
    if input == "inputfile.txt" {
        isFile = true
    } else {
        if !(IsIpv4Net(input) || IsIpv6Net(input) || strings.Count(input, ".") > 0) { 
            fmt.Println(Banner)
            return
        }
    }

    if isFile {
        file, err := os.Open(input)
        
       if err != nil {
          fmt.Println("failed opening file: %s", err)
            return
       }
        
        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)

        for scanner.Scan() {
            txtlines := scanner.Text()
            if len(txtlines) == 0 { continue }
            if !(IsIpv4Net(txtlines) || IsIpv6Net(txtlines) || strings.Count(txtlines, ".") > 0) {
                fmt.Println("Error: line [", errline, "] IP/Domain: [", txtlines, "] - check if the format is correct")
                errline++
                continue
            }
        
            result, err := whois.Whois(txtlines)
            if err != nil {
                fmt.Println ("Error: ", err)
                return
            }
            
            var xsip, xaddr, xOrg, xUpdated, xCountry   string
            xsip = txtlines
            
            if IsIpv4Net(txtlines) || IsIpv6Net (txtlines) {
                addr, _ := net.LookupAddr(txtlines)
                xaddr = addr[0]

                scanner := bufio.NewScanner(strings.NewReader(result))
                
                for scanner.Scan() {
                    linex := scanner.Text()
                    if !(strings.HasPrefix(linex,"#") || len(linex)==0 || strings.HasPrefix(linex,"Comment:")) {
                        if (strings.HasPrefix(linex,"Organization")) {
                            trimmedString := strings.TrimSpace(linex)
                            parts := strings.Split(trimmedString, ":")
                            trimmedString = strings.TrimSpace(parts[1])
                            xOrg = trimmedString
                        }
                        if (strings.HasPrefix(linex,"Updated")) {
                            trimmedString := strings.TrimSpace(linex)
                            parts := strings.Split(trimmedString, ":")
                            trimmedString = strings.TrimSpace(parts[1])
                            xUpdated = trimmedString
                        }
                        if (strings.HasPrefix(linex,"Country")) {
                            trimmedString := strings.TrimSpace(linex)
                            parts := strings.Split(trimmedString, ":")
                            trimmedString = strings.TrimSpace(parts[1])
                            xCountry = trimmedString
                        }
                    }
                }
                fmt.Println(xsip + "," + xaddr + "," + xOrg + "," + xUpdated + "," + xCountry)
                errline++
                continue
            }

            resultP, err := whoisparser.Parse(result)
            if err != nil {
                fmt.Println ("Error: ", err)
                return
            }

            var ddom, dip, dRegName, dUpdated, dCountry   string
            
            ip, _ := net.ResolveIPAddr("ip4", txtlines)

            ddom = txtlines
            dip  = ip.String()
        
            if len(resultP.Domain.UpdatedDate) > 0 { dUpdated = resultP.Domain.UpdatedDate }    
            if len(resultP.Registrant.Name) > 0 { dRegName = resultP.Registrant.Name }
            if len(resultP.Registrant.Country) > 0 { dCountry = resultP.Registrant.Country }
            
            fmt.Println(ddom + "," + dip + "," + dRegName + "," + dUpdated + "," + dCountry)
            errline++
        }
        file.Close()
        return
    }

    result, err := whois.Whois(input)
    if err != nil {
        fmt.Println ("Error: ", err)
        return
    }
    
    if IsIpv4Net(input) || IsIpv6Net (input) {
        addr, _ := net.LookupAddr(input)
    
        scanner := bufio.NewScanner(strings.NewReader(result))
        for scanner.Scan() {
            linex := scanner.Text()
            if !(strings.HasPrefix(linex,"#") || len(linex)==0 || strings.HasPrefix(linex,"Comment:")) { 
                switch Switch {
                case "C":
                    if (strings.HasPrefix(linex,"CIDR")) { fmt.Println(linex) }
                case "R":
                    //Do nothing, the Switch below will take care the intended output
                case "NIL":
                    fmt.Println(linex)
                default:
                    fmt.Println("Unrecognized switch!")
                }
            }
        }
        switch Switch {
            case "R":
                fmt.Println("Cannonical Name: ", addr)
            case "NIL":
                fmt.Println("\nCannonical Name: ", addr)
        }
        fmt.Println("")
        return
    }

    resultP, err := whoisparser.Parse(result)
    if err != nil {
        fmt.Println ("Error: ", err)
        return
    }
    
    ip, _ := net.ResolveIPAddr("ip4", input)
    fmt.Println("Resolved IP: \t\t\t", ip)
        
    switch Switch {
    case "S":
        fmt.Println("Status: \t\t\t", resultP.Domain.Status)        
    case "v":
        fmt.Println("Status: \t\t\t", resultP.Domain.Status)
        if resultP.Domain.DNSSec { fmt.Println("DNSSec: \t\t\t", resultP.Domain.DNSSec) }
        if len(resultP.Domain.CreatedDate) > 0 { fmt.Println("Created: \t\t\t", resultP.Domain.CreatedDate) }
        if len(resultP.Domain.UpdatedDate) > 0 { fmt.Println("Updated: \t\t\t", resultP.Domain.UpdatedDate) }    
        if len(resultP.Domain.ExpirationDate) > 0 { fmt.Println("Expiration: \t\t\t", resultP.Domain.ExpirationDate) }
        if len(resultP.Domain.ID) > 0 { fmt.Println("ID: \t\t\t\t", resultP.Domain.ID) }
        if len(resultP.Domain.NameServers) > 0 { fmt.Println("Name Servers: \t\t\t", resultP.Domain.NameServers) }
        if len(resultP.Domain.WhoisServer) > 0 { fmt.Println("Whois Server: \t\t\t", resultP.Domain.WhoisServer) }
        if len(resultP.Registrar.Name) > 0 { fmt.Println("Registrar Name: \t\t", resultP.Registrar.Name) }
        if len(resultP.Registrant.Name) > 0 { fmt.Println("Registrant Name: \t\t", resultP.Registrant.Name) }
        if len(resultP.Registrant.Email) > 0 { fmt.Println("Registrant Email: \t\t", resultP.Registrant.Email) }
        if len(resultP.Registrant.Phone) > 0 { fmt.Println("Registrant Phone: \t\t", resultP.Registrant.Phone) }
        if len(resultP.Registrant.Organization) > 0 { fmt.Println("Registrant Organization: \t", resultP.Registrant.Organization) }
        if len(resultP.Registrant.Street) > 0 { fmt.Println("Registrant Address: \t\t", resultP.Registrant.Street) }
        if len(resultP.Registrant.City) > 0 { fmt.Println("Registrant City: \t\t", resultP.Registrant.City) }
        if len(resultP.Registrant.Province) > 0 { fmt.Println("Registrant Province: \t\t", resultP.Registrant.Province) }    
        if len(resultP.Registrant.Country) > 0 { fmt.Println("Registrant Country: \t\t", resultP.Registrant.Country) }
        if len(resultP.Technical.PostalCode) > 0 { fmt.Println("Registrant Postal Code: \t", resultP.Registrant.PostalCode) }
        if len(resultP.Technical.Name) > 0 { fmt.Println("Technical Name: \t\t", resultP.Technical.Name) }
        if len(resultP.Technical.Email) > 0 { fmt.Println("Technical Email: \t\t", resultP.Technical.Email) }
        if len(resultP.Technical.Phone) > 0 { fmt.Println("Technical Phone: \t\t", resultP.Technical.Phone) }  
        if len(resultP.Administrative.Name) > 0 { fmt.Println("Administrative Name: \t\t", resultP.Administrative.Name) }
        if len(resultP.Administrative.Email) > 0 { fmt.Println("Administrative Email: \t\t", resultP.Administrative.Email) }
    case "N":
        if len(resultP.Registrant.Name) > 0 { fmt.Println("Registrant Name: \t\t", resultP.Registrant.Name) }
        if len(resultP.Registrant.Organization) > 0 { fmt.Println("Registrant Organization: \t", resultP.Registrant.Organization) }
        if len(resultP.Registrant.Country) > 0 { fmt.Println("Registrant Country: \t\t", resultP.Registrant.Country) }
    case "T":
        if len(resultP.Domain.CreatedDate) > 0 { fmt.Println("Created: \t\t\t", resultP.Domain.CreatedDate) }
        if len(resultP.Domain.UpdatedDate) > 0 { fmt.Println("Updated: \t\t\t", resultP.Domain.UpdatedDate) }    
        if len(resultP.Domain.ExpirationDate) > 0 { fmt.Println("Expiration: \t\t\t", resultP.Domain.ExpirationDate) }
    case "NIL":
        if resultP.Domain.DNSSec { fmt.Println("DNSSec: \t\t\t", resultP.Domain.DNSSec) }
        if len(resultP.Domain.CreatedDate) > 0 { fmt.Println("Created: \t\t\t", resultP.Domain.CreatedDate) }
        if len(resultP.Domain.UpdatedDate) > 0 { fmt.Println("Updated: \t\t\t", resultP.Domain.UpdatedDate) }    
        if len(resultP.Domain.ExpirationDate) > 0 { fmt.Println("Expiration: \t\t\t", resultP.Domain.ExpirationDate) }
        if len(resultP.Domain.NameServers) > 0 { fmt.Println("Name Servers: \t\t\t", resultP.Domain.NameServers) }
        if len(resultP.Domain.WhoisServer) > 0 { fmt.Println("Whois Server: \t\t\t", resultP.Domain.WhoisServer) }
        if len(resultP.Registrar.Name) > 0 { fmt.Println("Registrar Name: \t\t", resultP.Registrar.Name) }
        if len(resultP.Registrant.Name) > 0 { fmt.Println("Registrant Name: \t\t", resultP.Registrant.Name) }
        if len(resultP.Registrant.Organization) > 0 { fmt.Println("Registrant Organization: \t", resultP.Registrant.Organization) }
        if len(resultP.Registrant.Country) > 0 { fmt.Println("Registrant Country: \t\t", resultP.Registrant.Country) }        
    default:
        fmt.Println("Unrecognized switch!")
    }
}