# goview

PoC for embedding react app into native binary `~ 4MB` with webview.

### Usage
```bash
# Run the app
make run

# Build the binary
make build

```
```
$$  8:01PM [~/Workspace/projects/goview] (main) ✔ ls -lah ./goview
-rwxrwxr-x 1 snehesh snehesh 7.6M Jul  8 20:01 ./goview
$$  8:01PM [~/Workspace/projects/goview] (main) ✔ upx goview
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2020
UPX 3.96        Markus Oberhumer, Laszlo Molnar & John Reiser   Jan 23rd 2020

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   7876784 ->   4135080   52.50%   linux/amd64   goview

Packed 1 file.
$$  8:01PM [~/Workspace/projects/goview] (main) ✔ ls -lah ./goview
-rwxrwxr-x 1 snehesh snehesh 4.0M Jul  8 20:01 ./goview
```


### Screenshots
![Screenshot from 2022-07-08 19-19-01](https://user-images.githubusercontent.com/5787031/178083240-cefe8c41-2b5a-4208-8dba-d00dba1340ee.png)
