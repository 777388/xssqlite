# xssqlite
Find Subdomains, Grab All URLS for each, parse for fields, then XSS and SQL test it

Requirements: Go, Python, Massurl, GAU, Dalfox, SQLmap, Sublist3r

https://github.com/theFr1nge/massurl

https://github.com/lc/gau

https://github.com/hahwul/dalfox

https://github.com/sqlmapproject/sqlmap

https://github.com/aboul3la/Sublist3r



Usage: python xssqlite.py site.txt siteurl.txt domain.com

where site.txt is a list of urls you've gathered
and siteurls.txt is the output file
Will also output a second file of parsed fields
followed by a xss and sql readout which you have to pay attention while its running due to overfilling the terminal page.


Keep GAU, wsl, dalfox, sqlmap, sublist3r and massurl in the same folder unless you're willing to deal with path variables, I placed GAU and WSL in there already

Planned projectory: Spider that crawls Hackerone and Bugcrowd for in scope urls to feed the program.
