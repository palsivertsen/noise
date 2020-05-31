# Usage

```go
n := noise.Smooth{
	Size: 6000,
	Rand: rand.New(rand.NewSource(time.Now().UnixNano())), // Optional
}

for {
	fmt.Println(n.Next())
	time.Sleep(time.Millisecond * 10)
}
```

![smooth-noise](https://user-images.githubusercontent.com/11890143/83355149-d6d83100-a35d-11ea-8099-2087281bb5ad.gif)
