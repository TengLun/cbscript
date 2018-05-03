# cbscript
This is a simple program to add publishers, ip addresses, and device ids to your custom blacklist.

## Usage

Clone and build into an application. Commands and flags are available by typing:
```
cbscript --help
```

The following commands are available:
```
# Add the csv to the blacklist
cbscript add

# Remove the csv from the blacklist
cbscript remove

# Update the entires on the csv
cbscript update

```

Note the following flags are required:
```
# Always include your api key and file name
cbscript --api_key="YOUR API KEY" --file="YOUR FILE NAME" add

```

## Summary
This is a simple script to add entries to your blacklist. This project will eventually be merged with this project:
[Kochava Fraud API Library](https://github.com/TengLun/kfapi "kfAPI")
