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
var Ops bool


func IsIpv4Net(host string) bool {
   return net.ParseIP(host) != nil
}

func IsIpv6Net(host string) bool {
   return net.ParseIP(host) != nil
}

func main() { 
    defer func() {
        if r := recover(); r != nil {
            if !Ops {
                Banner := "whois-check v1.0\n\n"
                Banner = Banner + "Last Update: 10 Apr 2024, Alex Yang (https://linkedin.com/in/4yang)\n"
                Banner = Banner + "Usage: whois-check [ipv4 | ipv6 | domain.com]"
                fmt.Println(Banner)
            }
        }
    }()
    
    input := os.Args[1]
    if !(strings.Contains(input, ".") || strings.Contains(input, ":")) {
        fmt.Println("Usage: whois-check [ipv4 | ipv6 | domain.com]")
        return
    }
    
    result, err := whois.Whois(input)
    if err != nil {
        fmt.Println ("Error: ", err)
        return
    }
    
    if IsIpv4Net(input) || IsIpv6Net (input) {
        addr, _ := net.LookupAddr(input)
        //fmt.Println(result)
        
        scanner := bufio.NewScanner(strings.NewReader(result))
        for scanner.Scan() {
            linex := scanner.Text()
            //fmt.Println("[" + linex + "]")
            if !(strings.HasPrefix(linex,"#") || len(linex)==0 || strings.HasPrefix(linex,"Comment:")) { fmt.Println(linex) }
        }
    
        fmt.Println("\nReverse PTR Record: \t", addr)
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

    fmt.Println("Status: \t\t\t", resultP.Domain.Status)
    Ops = true
    if len(resultP.Domain.CreatedDate) > 0 { fmt.Println("Created: \t\t\t", resultP.Domain.CreatedDate) }
    if len(resultP.Domain.ExpirationDate) > 0 { fmt.Println("Expiration: \t\t\t", resultP.Domain.ExpirationDate) }
    if len(resultP.Domain.NameServers) > 0 { fmt.Println("Name Server: \t\t\t", resultP.Domain.NameServers) }
    if len(resultP.Domain.WhoisServer) > 0 { fmt.Println("Whois Server: \t\t\t", resultP.Domain.WhoisServer) }
    //if len(resultP.Domain.DNSSEC) > 0 { fmt.Println("DNSSEC: \t\t\t", resultP.Domain.DNSSEC) }
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
}

/*
type Contact struct {
    ID           string `json:"id,omitempty"`
    Name         string `json:"name,omitempty"`
    Organization string `json:"organization,omitempty"`
    Street       string `json:"street,omitempty"`
    City         string `json:"city,omitempty"`
    Province     string `json:"province,omitempty"`
    PostalCode   string `json:"postal_code,omitempty"`
    Country      string `json:"country,omitempty"`
    Phone        string `json:"phone,omitempty"`
    PhoneExt     string `json:"phone_ext,omitempty"`
    Fax          string `json:"fax,omitempty"`
    FaxExt       string `json:"fax_ext,omitempty"`
    Email        string `json:"email,omitempty"`
    ReferralURL  string `json:"referral_url,omitempty"`
}

type Domain struct {
    ID             string `json:"id,omitempty"`
    Domain         string `json:"domain,omitempty"`
    Name           string `json:"name,omitempty"`
    Extension      string `json:"extension,omitempty"`
    Status         string `json:"status,omitempty"`
    DNSSEC         string `json:"dnssec,omitempty"`
    WhoisServer    string `json:"whois_server,omitempty"`
    NameServers    string `json:"name_servers,omitempty"`
    CreatedDate    string `json:"created_date,omitempty"`
    UpdatedDate    string `json:"updated_date,omitempty"`
    ExpirationDate string `json:"expiration_date,omitempty"`
}

type WhoisInfo struct {
    Domain         *Domain  `json:"domain,omitempty"`
    Registrar      *Contact `json:"registrar,omitempty"`
    Registrant     *Contact `json:"registrant,omitempty"`
    Administrative *Contact `json:"administrative,omitempty"`
    Technical      *Contact `json:"technical,omitempty"`
    Billing        *Contact `json:"billing,omitempty"`
}
*/