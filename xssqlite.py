import os
import sys
f = ("python Sublist3r/sublist3r.py -d "+sys.argv[3]+ " -o "+sys.argv[1]+" && go run crawl.go "+sys.argv[1]+" "+sys.argv[2]+" | python xssql2.py ;"+" go run xssql.go 2"+sys.argv[2]+" | python xssql2.py")
os.system(f)