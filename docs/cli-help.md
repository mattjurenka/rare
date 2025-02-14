% rare(8) Aggregate and display information parsed from text files using
	regex and a simple handlebars-like expressions.

	Run "rare docs overview" or go to https://rare.zdyn.net for more information
	
	https://github.com/zix99/rare

% 

# NAME

rare - A fast regex parser, extractor and realtime aggregator

# SYNOPSIS

rare

```
[--color]
[--help|-h]
[--nocolor|--nc]
[--noformat|--nf]
[--notrim]
[--nounicode|--nu]
[--profile]=[value]
[--version|-v]
```

**Usage**:

```
rare [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--color**: Force-enable color output

**--help, -h**: show help

**--nocolor, --nc**: Disables color output

**--noformat, --nf**: Disable number formatting

**--notrim**: By default, rare will trim output text for in-place updates. Setting this flag will disable that

**--nounicode, --nu**: Disable usage of unicode characters

**--profile**="": Write application profiling information as part of execution. Specify base-name

**--version, -v**: print the version


# COMMANDS

## filter, f

Filter incoming results with search criteria, and output raw matches

**--batch**="": Specifies io batching size. Set to 1 for immediate input (default: 1000)

**--extract, -e**="": Expression that will generate the key to group by. Specify multiple times for multi-dimensions or use {$} helper (default: {0})

**--follow, -f**: Read appended data as file grows

**--gunzip, -z**: Attempt to decompress file when reading

**--ignore, -i**="": Ignore a match given a truthy expression (Can have multiple)

**--ignore-case, -I**: Augment regex to be case insensitive

**--line, -l**: Output source file and line number

**--match, -m**="": Regex to create match groups to summarize on (default: .*)

**--poll**: When following a file, poll for changes rather than using inotify

**--posix, -p**: Compile regex as against posix standard

**--readers, --wr**="": Sets the number of concurrent readers (Infinite when -f) (default: 3)

**--recursive, -R**: Recursively walk a non-globbing path and search for plain-files

**--reopen, -F**: Same as -f, but will reopen recreated files

**--tail, -t**: When following a file, navigate to the end of the file to skip existing content

**--workers, -w**="": Set number of data processors (default: 3)

**-n, --num**="": Print the first NUM of lines seen (Not necessarily in-order) (default: 0)

## histogram, h, histo

Summarize results by extracting them to a histogram

**--all, -a**: After summarization is complete, print all histogram buckets

**--atleast**="": Only show results if there are at least this many samples (default: 0)

**--bars, -b**: Display bars as part of histogram

**--batch**="": Specifies io batching size. Set to 1 for immediate input (default: 1000)

**--extra, -x**: Alias for -b --percentage

**--extract, -e**="": Expression that will generate the key to group by. Specify multiple times for multi-dimensions or use {$} helper (default: {0})

**--follow, -f**: Read appended data as file grows

**--gunzip, -z**: Attempt to decompress file when reading

**--ignore, -i**="": Ignore a match given a truthy expression (Can have multiple)

**--ignore-case, -I**: Augment regex to be case insensitive

**--match, -m**="": Regex to create match groups to summarize on (default: .*)

**--num, -n**="": Number of elements to display (default: 5)

**--percentage**: Display percentage of total next to the value

**--poll**: When following a file, poll for changes rather than using inotify

**--posix, -p**: Compile regex as against posix standard

**--readers, --wr**="": Sets the number of concurrent readers (Infinite when -f) (default: 3)

**--recursive, -R**: Recursively walk a non-globbing path and search for plain-files

**--reopen, -F**: Same as -f, but will reopen recreated files

**--reverse**: Reverses the display sort-order

**--sortkey, --sk**: Sort by key, rather than value

**--tail, -t**: When following a file, navigate to the end of the file to skip existing content

**--workers, -w**="": Set number of data processors (default: 3)

## bargraph, b, bar, bars

Create a bargraph of the given 1 or 2 dimension data

**--batch**="": Specifies io batching size. Set to 1 for immediate input (default: 1000)

**--extract, -e**="": Expression that will generate the key to group by. Specify multiple times for multi-dimensions or use {$} helper (default: {0})

**--follow, -f**: Read appended data as file grows

**--gunzip, -z**: Attempt to decompress file when reading

**--ignore, -i**="": Ignore a match given a truthy expression (Can have multiple)

**--ignore-case, -I**: Augment regex to be case insensitive

**--match, -m**="": Regex to create match groups to summarize on (default: .*)

**--poll**: When following a file, poll for changes rather than using inotify

**--posix, -p**: Compile regex as against posix standard

**--readers, --wr**="": Sets the number of concurrent readers (Infinite when -f) (default: 3)

**--recursive, -R**: Recursively walk a non-globbing path and search for plain-files

**--reopen, -F**: Same as -f, but will reopen recreated files

**--reverse**: Reverses the display sort-order

**--stacked, -s**: Display bargraph as stacked

**--tail, -t**: When following a file, navigate to the end of the file to skip existing content

**--workers, -w**="": Set number of data processors (default: 3)

## analyze, a

Numerical analysis on a set of filtered data

**--batch**="": Specifies io batching size. Set to 1 for immediate input (default: 1000)

**--extra, -x**: Displays extra analysis on the data (Requires more memory and cpu)

**--extract, -e**="": Expression that will generate the key to group by. Specify multiple times for multi-dimensions or use {$} helper (default: {0})

**--follow, -f**: Read appended data as file grows

**--gunzip, -z**: Attempt to decompress file when reading

**--ignore, -i**="": Ignore a match given a truthy expression (Can have multiple)

**--ignore-case, -I**: Augment regex to be case insensitive

**--match, -m**="": Regex to create match groups to summarize on (default: .*)

**--poll**: When following a file, poll for changes rather than using inotify

**--posix, -p**: Compile regex as against posix standard

**--quantile, -q**="": Adds a quantile to the output set. Requires --extra (default: 90,99,99.9)

**--readers, --wr**="": Sets the number of concurrent readers (Infinite when -f) (default: 3)

**--recursive, -R**: Recursively walk a non-globbing path and search for plain-files

**--reopen, -F**: Same as -f, but will reopen recreated files

**--reverse, -r**: Reverses the numerical series when ordered-analysis takes place (eg Quantile)

**--tail, -t**: When following a file, navigate to the end of the file to skip existing content

**--workers, -w**="": Set number of data processors (default: 3)

## tabulate, t, table

Create a 2D summarizing table of extracted data

**--batch**="": Specifies io batching size. Set to 1 for immediate input (default: 1000)

**--cols**="": Number of columns to display (default: 10)

**--delim**="": Character to tabulate on. Use {$} helper by default (default:  )

**--extract, -e**="": Expression that will generate the key to group by. Specify multiple times for multi-dimensions or use {$} helper (default: {0})

**--follow, -f**: Read appended data as file grows

**--gunzip, -z**: Attempt to decompress file when reading

**--ignore, -i**="": Ignore a match given a truthy expression (Can have multiple)

**--ignore-case, -I**: Augment regex to be case insensitive

**--match, -m**="": Regex to create match groups to summarize on (default: .*)

**--num, -n**="": Number of elements to display (default: 20)

**--poll**: When following a file, poll for changes rather than using inotify

**--posix, -p**: Compile regex as against posix standard

**--readers, --wr**="": Sets the number of concurrent readers (Infinite when -f) (default: 3)

**--recursive, -R**: Recursively walk a non-globbing path and search for plain-files

**--reopen, -F**: Same as -f, but will reopen recreated files

**--sortkey, --sk**: Sort rows by key name rather than by values

**--tail, -t**: When following a file, navigate to the end of the file to skip existing content

**--workers, -w**="": Set number of data processors (default: 3)

## docs

Access detailed documentation

**--no-pager, -n**: Don't use pager to view documentation

## help, h

Shows a list of commands or help for one command
