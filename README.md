## Parlia simulation
+ Multi-Decree Paxos algorithm toy implementation.
+ Non-deterministic simulation with pseudo-random network delays and nodes deaths.

### Run Configuration

Amount of entities depends on line count in configs

Proxy: [file](./config/proxy_ports.txt)
+ Port for each proxy

Replica [file](./config/replica_config.txt)
- Port for each replica
- Run mode:
    + Trust - network delays
    + Faulty - network delays + node deaths

### Origin
+ [Simple](https://lamport.azurewebsites.net/pubs/paxos-simple.pdf)
+ [Greek](https://lamport.azurewebsites.net/pubs/lamport-paxos.pdf)
+ [Pseudo](https://pdos.csail.mit.edu/archive/6.824-2013/notes/paxos-code.html)

