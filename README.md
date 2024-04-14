# whois-check

Return domain registrar info and resolved IP (or Cannonical Name) using WHOIS DB and Li Kexian's whois library

The binary for MacOS (compiled on Sonoma 14.x) is included in this repository.

```
whois-check v2.0
Last Update: 14 Apr 2024, Alex Yang (https://linkedin.com/in/4yang)

Usage for Single IP query:
    whois-check [ipv4 | ipv6 | domain.com]

Optional_Switch for output format (FOR DOMAIN ONLY):
    S   Show only resolved IP & domain status
    v   Show verbose (more) information
    N   Show only resolved IP, registrant name, organization & country
    T   Show only created, updated and expiration date

Example:
   whois-check twitter.com
   whois-check google.com c
   whois-check youtube.com N
   whois-check netflix.com T

Optional_Switch for output format (FOR IPv4/v6 ONLY):
    C   Show only CIDR
    R   Show only Cannonical Name

Example:
   whois-check 20.231.239.246 C
   whois-check 142.251.12.94 R

Usage for Bulk IP query:
   whois-check inputfile.txt
   (the input filename must be STRICTLY inputfile.txt. Lines must only a IPv4/6 and domain name)

Example:
   whois-check input.txt


Usage Example:
    %  whois-check 20.231.239.246 C

Output:
    CIDR:           20.192.0.0/10

Example:
    % whois-check 142.251.12.94 R

Output:
    Reverse PTR Record:  [se-in-f94.1e100.net.]

Example:
    % whois-check google.com T

Output:
    Created:                   1997-09-15T04:00:00Z
    Updated:                   2019-09-09T15:39:04Z
    Expiration:                2028-09-14T04:00:00Z

Example:
    % whois-check twitter.com
   
Output:
    Resolved IP:               104.244.42.65
    Created:                   2000-01-21T16:28:17Z
    Updated:                   2024-01-17T06:10:05Z
    Expiration:                2025-01-21T16:28:17Z
    Name Servers:              [a.r06.twtrdns.net a.u06.twtrdns.net b.r06.twtrdns.net b.u06.twtrdns.net c.r06.twtrdns.net c.u06.twtrdns.net d.r06.twtrdns.net d.u06.twtrdns.net]
    Whois Server:              whois.corporatedomains.com
    Registrar Name:            CSC Corporate Domains, Inc.
    Registrant Name:           Twitter, Inc.
    Registrant Organization:   Twitter, Inc.
    Registrant Country:        US

Example:
    % whois-check 2607:f8b0:4003:c00::6a

Output:
    NetRange:       2607:F8B0:: - 2607:F8B0:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF
    CIDR:           2607:F8B0::/32
    NetName:        GOOGLE-IPV6
    NetHandle:      NET6-2607-F8B0-1
    Parent:         NET6-2600 (NET6-2600-1)
    NetType:        Direct Allocation
    OriginAS:       AS22577, AS15169
    Organization:   Google LLC (GOGL)
    RegDate:        2009-03-12
    Updated:        2012-02-24
    Ref:            https://rdap.arin.net/registry/ip/2607:F8B0::
    OrgName:        Google LLC
    OrgId:          GOGL
    Address:        1600 Amphitheatre Parkway
    City:           Mountain View
    StateProv:      CA
    PostalCode:     94043
    Country:        US
    RegDate:        2000-03-30
    Updated:        2019-10-31
    Ref:            https://rdap.arin.net/registry/entity/GOGL
    OrgTechHandle:  ZG39-ARIN
    OrgTechName:    Google LLC
    OrgTechPhone:   +1-650-253-0000 
    OrgTechEmail:   arin-contact@google.com
    OrgTechRef:     https://rdap.arin.net/registry/entity/ZG39-ARIN
    OrgAbuseHandle: ABUSE5250-ARIN
    OrgAbuseName:   Abuse
    OrgAbusePhone:  +1-650-253-0000 
    OrgAbuseEmail:  network-abuse@google.com
    OrgAbuseRef:    https://rdap.arin.net/registry/entity/ABUSE5250-ARIN
    % Query time: 1081 msec
    % WHEN: Wed Apr 10 23:13:33 +08 2024

    Reverse PTR Record: [nd-in-f106.1e100.net.]
```

# If you have any issue and need a little help

Please don't hesitate to DM me at **Linkedin** OR open an issue.

https://linkedin.com/in/4yang

# To contribute

Please make a PR to help improve this tool :)