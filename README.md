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

Killer will exit with status code `0` if it was able to successfully stop a process. If
the pid of a non-existent process is given, this will be treated as a success.

It will exit with status code `1` if a signal could not be sent (due to permission errors)
or if the process did not stop in time.

## Motivation

Why does this exist? Why build a glorified version of the `kill` command? Some backstory...

The Java Virtual Machine (JVM) has a flag to run a particular command when it encounters
an out of memory condition (OOM). Some distributed software, like ZooKeeper, recommends
that you set the JVM to kill itself when it encounters an OOM error. The command looks
something like...

```
java -XX:OnOutOfMemoryError='kill -9 %p' ...
```

The motivation for this is that when an `OutOfMemoryError` gets raised, you don't know
what random thread got killed and you can't be sure your application is still healthy.

This works well for ZooKeeper since it's a mature and battle tested piece of software.
You can be confident that your data won't be corrupted or lost when you `kill -9` ZooKeeper.

However, not all software is this robust. It'd be nice to give the application a chance
to shutdown gracefully before resorting to `kill -9` when the application encounters
and OOM error.

Hence: the behavior of `killer`.

```
java -XX:OnOutOfMemoryError='killer %p' ...
```

## Building

TBD
