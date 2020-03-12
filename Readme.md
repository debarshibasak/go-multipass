## go-mulitpass client

Golang client SDK to interact with canonical's multipass.

#### To launch an instance
```
instance, err := multipass.Launch(&multipass.LaunchReq{
    CPU:           2,
    Name: "instanceName", 
})
if err != nil {
    t.Fatal(err)
}
```

#### To delete an instance

```
if err := multipass.Delete(&multipass.DeleteRequest{Name:instance.Name}); err != nil {
    t.Fatal(err)
}
```

#### To get information an instance

```
instanceInfo, err := multipass.Info(&multipass.InfoRequest{Name:instance.Name})
if err != nil {
    t.Fatal(err)
}
```


#### To execute remote commands

```
err = multipass.Exec(&multipass.ExecRequest{Name:instanceInfo.Name, Command:"ls"})
if err != nil {
    t.Fatal(err)
}

```

#### Roadmap

- Cover all the commandline parameters of multipass
- More documentation and examples to be added