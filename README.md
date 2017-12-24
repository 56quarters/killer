# Killer

Repeatedly try to stop a process with `SIGTERM` and eventually `SIGKILL`.

## Usage

```
killer [--interval int] [--timeout int] [--disable-kill] [PID]
```

The `killer` command will repeatedly send a process (`PID`) the `SIGTERM` signal until
it stops and will finally send a `SIGKILL` if the process does not stop within the given
timeout.

`--interval int` How long to wait in between each `SIGTERM`, in seconds

`--timeout int` How long to wait total before sending a `SIGKILL`, in seconds

`--disable-kill` Don't send a final `SIGKILL` after repeatedly sending `SIGTERM`

## Motivation

Why does this exist? Why build a glorified version of the `kill` command? Some backstory...

The Java Virtual Machine (JVM) has a flag to run a particular command when it encounters
and out of memory condition (OOM). Some software, like ZooKeeper, recommends that you set
the JVM to kill itself when it encounters an OOM error. The command looks something like

```
java -XX:OnOutOfMemoryError='kill -9 %p' ...
```

The motivation for this is that when an `OutOfMemoryError` gets raised, you don't know
what random thread got killed and you can't be sure your application is still healthy.

This works well for ZooKeeper since it's a mature and battle tested piece of software.
There's not a lot of danger of data corruption or loss when you `kill -9` it.

However, not all software is this robust. It'd be nice to give the application a chance
to shutdown gracefully before resorting to `kill -9` when the application encounters
and OOM error.

Hence: the behavior of `killer`.

## Building

TBD
