- What do we have in the directory

- Plugin unittest installed 
```
helm plugin install https://github.com/helm-unittest/helm-unittest
```

$ helm unittest .

- Run test by test
    - First check only if the Deployment exist


How and when the tests are used
- it can be used in pipelines 
- when a base chart is crated
- when an additional chart is used for one specific depratment let's say sales, the base chart source