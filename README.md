# goaccess-json-parser
- program written in golang to parse json output of goaccess tool into a smaller json file for website analytics tracking.
- report.json is a sample output from the goaccess tool ran on an nginx access log
- running the parseLog.go program should result in a json file output called analytics.json
    - this file should contain following analytics data:
        - number of unique visitors
        - regions with visitors count followed by vistors percentage
        - operating systems with visitors count followed by visitors percentage
