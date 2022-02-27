# xssqlite
Grab All URLS, parse for fields, then XSS and SQL test it

Requirements: Python, Massurl, GAU, Dalfox, SQLmap



Usage: python xssqlite.py site.txt siteurl.txt
where site.txt is a list of urls you've gathered
and siteurls.txt is the output file
Will also output a second file of parsed fields
followed by a xss and sql readout
