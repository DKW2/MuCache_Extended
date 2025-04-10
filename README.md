# Mucache_Extended

This is a capstone project to analyze and look deeper into the microservice caching framework: MuCache. It focuses on reproducing the results of the paper (to confirm their findings) as well as explore how different caching strategies can affect the performance of the cache. This is inspired by the fact that the MuCache paper uses a simple key-value caching/invalidation strategy and did not explore any further. Thus, by exploring more advanced strategies, we can see how different caching strategies can affect the performance of microservices.

The paper describing the entire MuCache framework and experiments is in this NSDI 2024 paper:
[Mucache: a General Framework for Caching in Microservice Graphs](https://www.usenix.org/conference/nsdi24/presentation/zhang-haoran).

## How to run the code and reproduce the results 
See [scripts/README.md](scripts/README.md) for a detailed description of 
how to set up our code, run our experiments, and reproduce our results. *Will update for more clarity in this repo*

## Citing the Paper

```
@inproceedings{zhang24mucache,
 author = {Haoran Zhang and Konstantinos Kallas and Spyros Pavlatos and Rajeev Alur and Sebastian Angel and Vincent Liu},
 title = {Mucache: a General Framework for Caching in Microservice Graphs},
 booktitle = {USENIX Symposium on Networked Systems Design and Implementation (NSDI)},
 month = {April},
 year = {2024}
}
```
