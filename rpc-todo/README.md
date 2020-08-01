# RPC Server 

We will use the golang `net/rpc` package to expose the golang methods as RPC endpoints. 
> A RPC is when a computer program causes a procedure (subroutine) to execute in a different address
> space (commonly on another computer on a shared network), which is coded as if it were a normal
> (local) procedure call, without the programmer explicitly coding the details for the remote
> interaction. 

The `net/rpc` package in golang has the following requirements for exporting a func as an RPC: 
* The method type is exported. 
* The method is exported 
* The method has two arguments, both exported (or builtin types) 
* The method's second argument must be a pointer. 
* The method return type has to be an error. 

 
