# Downtime Monitor

This commandline tool queries the webservice reachable at the specified address, and starts a downtime timer once it becomes unreachable.

Once the service is answering with a status code of 200 again,
the timer is stopped and the downtime calculated.
