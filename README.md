# Splitsies

Splitsies is a CLI Tool that is to be used to help break a large CSV into smaller CSV files, each generated file to have their names be derived from a column value. 

### Usage

Source File:

```
date,user_id,action
2019-11-01,1001,click
2019-11-01,1002,click
2019-11-02,1003,click
2019-11-03,1004,click
2019-11-02,1004,click
```

Lets run splitsies on this
```sh
splitsies --header=true --out-file-prefix=day_ --out-file-col-index=0 --out-file-col-max-length=10 --file=source.csv --out-dir=out/
```

Will generate three files following files:
```
Filename: 2019-11-01.csv
2019-11-01,1001,click
2019-11-01,1002,click
```

```
Filename: 2019-11-02.csv
2019-11-02,1003,click
2019-11-02,1004,click
```

```
Filename: 2019-11-03.csv
2019-11-03,1004,click
```

The output files do not have the header.