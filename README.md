# xssqlite
Grab All URLS, parse for fields, then XSS and SQL test it

Requirements: Go, Python, Massurl, GAU, Dalfox, SQLmap

https://github.com/theFr1nge/massurl

https://github.com/lc/gau

https://github.com/hahwul/dalfox

https://github.com/sqlmapproject/sqlmap



Usage: python xssqlite.py site.txt siteurl.txt

where site.txt is a list of urls you've gathered
and siteurls.txt is the output file
Will also output a second file of parsed fields
followed by a xss and sql readout which you have to pay attention while its running due to overfilling the terminal page.


Keep GAU, wsl, dalfox, sqlmap and massurl in the same folder unless you're willing to deal with path variables, I placed GAU and WSL in there already

Planned projectory: Spider that crawls Hackerone and Bugcrowd for in scope urls to feed the program.
