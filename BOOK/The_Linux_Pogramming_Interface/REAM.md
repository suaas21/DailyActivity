# An overview of The Linux Programming Interface

## Fundamental Concepts

### Necessary Information

**Kernel:**

- The term kernel is often used as a synonym for the central software that manages and allocates computer resources (i.e., the CPU, RAM, and devices).

- The Linux kernel executable typically resides at the pathname `/boot/vmlinuz`, or something similar. The derivation of this filename is historical. On early UNIX implementations, the kernel was called unix . Later UNIX implementations, which implemented virtual memory, renamed the kernel as `vmunix` . On Linux, the filename mirrors the system name, with the z replacing the final x to signify that the kernel is a compressed executable.

- The kernel performs the following tasks: Process scheduling, Memory management, Provision of a file system, Creation and termination of processes, Access to devices, Networking and Provision of a system call application programming interface (API): _Processes can request the kernel to perform various tasks using kernel entry points known as system calls. The Linux system call API is the primary topic of this book._

- `Kernel mode and user mode:` Modern processor architectures typically allow the CPU to operate in at least two different modes: user mode and kernel mode (sometimes also referred to as supervisor mode).

**The Shell:**

- A shell is a special-purpose program designed to read commands typed by a user and execute appropriate programs in response to those commands. Such a program is sometimes known as a command interpreter. Whereas on some operating systems the command interpreter is an integral part of the kernel, on UNIX systems. Various type of shell:
  - `Bourne shell (sh):` This is the oldest of the widely used shells, and was written by Steve Bourne. It was the standard shell for Seventh Edition UNIX. The Bourne shell contains many of the features familiar in all shells: I/O redirection, pipelines filename generation (globbing), variables, manipulation of environment variables, command substitution, background command execution, and functions. All later UNIX implementations include the Bourne shell in addition to any other shells they might provide.
  - `C shell (csh):` This shell was written by Bill Joy at the University of California at Berkeley. The name derives from the resemblance of many of the flow-control constructs of this shell to those of the C programming language. The C shell provided several useful interactive features unavailable in the Bourne shell, including command history, command-line editing, job control, and aliases. The C shell was not backward compatible with the Bourne shell. Although the standard interactive shell on BSD was the C shell, shell scripts (described in a moment) were usually written for the Bourne shell, so as to be portable across all UNIX implementations.
  - `Korn shell (ksh):` This shell was written as the successor to the Bourne shell by David Korn at AT&T Bell Laboratories. While maintaining backward compatibility with the Bourne shell, it also incorporated interactive features similar to those provided by the C shell.
  - `Bourne again shell (bash):` This shell is the GNU project’s reimplementation of the Bourne shell. It supplies interactive features similar to those available in the C and Korn shells.

**Users and Groups:**

- __Users:__ Every user of the system has a unique login name (`username`) and a corresponding numeric user `ID` (`UID`). For each user, these are defined by a line in the system password file, `/etc/passwd` , which includes the following additional information:
  - `Group ID:` the numeric `group ID` of the first of the groups of which the user is a member.
  - `Home directory:` the initial directory into which the user is placed after logging in.
  - `Login shell`: the name of the program to be executed to interpret user commands.
- __Groups:__ For administrative purposes—in particular, for controlling access to files and other system resources—it is useful to organize users into groups. BSD allowed a user to simultaneously belong to multiple groups. Each group is identified by a single line in the system group file, `/etc/group` , which includes the following information:
  - `Group name:` the (unique) name of the group.
  - `Group ID (GID):` the numeric ID associated with this group.
  - `User list:` a comma-separated list of login names of users who are members of this group.
- __Superuser:__ One user, known as the superuser, has special privileges within the system. The superuser account has user ID 0, and normally has the login name root. On typical UNIX systems, the superuser bypasses all permission checks in the system.

**Single Directory Hierarchy, Directories, Links, and Files:**

- __Directory Hierarchy:__  The kernel maintains a single hierarchical directory structure to organize all files in the system. (This contrasts with operating systems such as Microsoft Windows, where each disk device has its own directory hierarchy.) At the base of this hierarchy is the root directory, named `/ (slash)`. All files and directories are children or further removed descendants of the root directory.

<figure align="center">
  <img alt="Stash Batch Backup Flow" src="/images/dh.png">
  <figcaption align="center">Fig: Subset of the Linux single directory hierarchy</figcaption>
</figure>

- __File types:__ Within the file system, each file is marked with a type, indicating what kind of file it is. One of these file types denotes ordinary data files, which are usually called regular or plain files to distinguish them from other file types. These other file types include devices, pipes, sockets, directories, and symbolic links. The term file is commonly used to denote a file of any type, not just a regular file.

- __Directories and links:__  From the above fig. you can find this information:
  - All the nodes of the tree, except the leaves, denote directory names. A directory node contains information about the files and directories just beneath it.
  - The directory corresponding to the root of the tree is called the root directory(`/`). To identify a specific file, the process uses a pathname, which consists of slashes alternating with a sequence of directory names that lead to the file. If the first item in the pathname is a slash, the pathname is said to be `absolute`, because its starting point is the root directory. Otherwise, if the first item is a directory name or filename, the pathname is said to be `relative`, because its starting point is the process’s current directory. While specifying filenames, the notations `“.`” and “`..`” are also used. They denote the current `working` directory and its `parent` directory, respectively. If the current working directory is the root directory, “`.`” and “`..`” coincide (`/.. equates to /`).
  - `Hard and Soft Links:`
      - `Hard Link:`A filename included in a directory is called a file hard link, or more simply, a link. The same file may have several links included in the same directory or in different ones, so it may have several filenames. linux command:
        ```console
        $ ln p1 p2
        ```
        is used to create a new hard link that has the pathname p2 for a file identified by the pathname p1.
      - `Soft Link:` To overcome hard link limitations, soft links (also called symbolic links) were introduced a long time ago. Linux command:
        ```console
        $ ln -s p1 p2
        ``` 
        You can find more details from here [difference between soft and hard link](https://www.ostechnix.com/explaining-soft-link-and-hard-link-in-linux-with-examples/)
- __File ownership and permissions:__ Each file has an associated user ID and group ID that define the owner of the file and the group to which it belongs. The ownership of a file is used to determine the access rights available to users of the file.Three permission bits may be set for each of these categories of user: `read`, `write` and `execute`.
- __File I/O Model:__ One of the distinguishing features of the I/O model on UNIX systems is the concept of universality of I/O. This means that the same system calls (open(), read(), write(), close(), and so on) are used to perform I/O on all types of files, including
devices.
- __File descriptors:__ The I/O system calls refer to open files using a file descriptor, a (usually small) non-negative integer. A file descriptor is typically obtained by a call to open(). Normally, a process inherits three open file descriptors when it is started by he shell: descriptor `0` is `standard input`, the file from which the process takes its input; descriptor `1` is `standard output`, the file to which the process writes its output; and descriptor `2` is `standard error`. more details example [here](https://www.computerhope.com/jargon/f/file-descriptor.htm).
- __The stdio library:__ To perform file I/O, C programs typically employ I/O functions contained in the standard C library. This set of functions, referred to as the stdio library, includes fopen(), fclose(), scanf(), printf(), fgets(), fputs(), and so on. The stdio functions are layered on top of the I/O system calls (open(), close(), read(), write(), and so on).

**Programs:** Programs normally exist in two forms. 1. source code like c code and 2. machine readable binary machine-language.
  - __filters:__ Reads its input from `stdin` and gives output to `stdout` example: `cat`, `grep`, `tr`, `sort`, `wc`, `sed`, and `awk`
  - __Command-line arguments:__ In C, programs can access the command-line arguments, the words that are supplied on the command line when the program is run. To access the command-line arguments, the main() function of the program is declared as follows:
    ``` 
     int main(int argc, char *argv[])
    ```
    The `argc` variable contains the total number of command-line arguments, and the individual arguments are available as strings pointed to by members of the array `argv`. The first of these strings, `argv[0]`, identifies the name of the program itself.
    
  **Process:**
   - __Process memory layout:__ A process is logically divided into the following parts, known as segments:
      - `Text:` the instructions of the program.
      - `Data:` the static variables used by the program.
      - `Heap:` an area from which programs can dynamically allocate extra memory.
      - `Stack:` a piece of memory that grows and shrinks as functions are called and return and that is used to allocate storage for local variables and function call linkage information.
  - __Process creation and program execution:__ A process can create a new process using the `fork()` system call. The process that calls `fork()` is referred to as the parent process, and the new process is referred to as the child process. The kernel creates the child process by making a duplicate of the parent process. Each process has a unique integer process identifier (`PID`). Each process also has a parent process identifier (`PPID`).
  - __The init process:__ When booting the system, the kernel creates a special process called init, the “parent of all processes,” which is derived from the program file `/sbin/init` . All processes on the system are created (using fork()) either by init or by one of its descendants. The init process always has the process `ID 1` and runs with superuser privileges. The init process can’t be killed (not even by the superuser), and it terminates only when the system is shut down. The main task of init is to create and monitor a range of processes required by a running system.
  - __Daemon processes:__
      - It is long-lived. A daemon process is often started at system boot and remains in existence until the system is shut down.
      - It runs in the background, and has no controlling terminal from which it can read input or to which it can write output.
      - example: `syslogd`, `httpd`.
  - __Interprocess Communication and Synchronization:__ Linux, like all modern UNIX implementations, provides a rich set of mechanisms for interprocess communication (IPC), including the following:
    - `signals`, which are used to indicate that an event has occurred; pipes (familiar to shell users as the `|` operator) and FIFOs, which can be used to transfer data between processes;
    - `sockets`, which can be used to transfer data from one process to another, either on the same host computer or on different hosts connected by a network;
    - `file locking`, which allows a process to lock regions of a file in order to prevent other processes from reading or updating the file contents;
    - `message queues`, which are used to exchange messages (packets of data) between processes;
    - `semaphores`, which are used to synchronize the actions of processes; and
    - `shared memory`, which allows two or more processes to share a piece of memory. When one process changes the contents of the shared memory, all of the other processes can immediately see the changes.
- **Signals:** Signals are sent to a process by the kernel, by another process (with suitable permissions), or by the process itself. For example, the kernel may send a signal to a process when one of the following occurs:
  - the user typed the interrupt character (usually Control-C) on the keyboard;
  - one of the process’s children has terminated;
  - a timer (alarm clock) set by the process has expired; or
  - the process attempted to access an invalid memory address.Within the shell, the `kill` command can be used to send a signal to a process. The `kill()` system call provides the same facility within programs. 