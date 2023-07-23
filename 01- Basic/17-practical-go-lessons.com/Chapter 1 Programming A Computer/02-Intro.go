/*

The four hardware components:
	#1 The memory unit (MU) where data and programs are stored.
	
	#2 The arithmetic and logic unit (ALU).
		Its role is to perform arithmetic and logical operations on data stored into the memory unit.
		This unit can perform, for instance, additions, incrementations, decrementations, which are called operations. 
		In general, each operation requires two operands and output a result.

	#3 The input and output unit (I/OU) will be in charge of loading data into the memory unit from an input device. 
		This unit also sends data from the memory unit to an output device.
		An input device is, for example, the touchpad of your computer, your keyboard, your mouse.
		An output device is, for instance, your monitor.

	#4 The control unit (CU) will receive instructions from programs and will control the activity of the other units. 


Memory
	A computer is composed of two types of memory :
		#1 The central memory
		#2 The auxiliary memory
	Two categories of memory exist:
		#1 Volatile
		#2 Non Volatile


The central memory:
	#1 RAM (Random Access Memory). *[This type of memory is volatile]*
		This type of storage requires electric power to persist data. 
		When you turn your computer off, the memory contained in this type of storage will be erased.
		The operating system and the programs you use will be loaded into this memory.

	#2 ROM (Read-Only Memory). *[This kind of memory is not volatile]*
		This is a memory that contains data necessary for the computer to run correctly. 
		It’s designed to be only readable and not updated by the system.
	

The auxiliary memory:
	This type of memory is not volatile.
	examples of auxiliary memory: Hard drives, USB keys, CD-ROM, DVD ...etc.
	# Read and writes to this type of memory is slow compared to the RAM.
	# Some hard drives sequentially access memory.
	# The system should respect a particular sequence.
	# Respecting this access sequence takes a longer time than a random access mode.
	# Note that some hard drives allow random memory access.


SSD hard drive:
	Hard drives, also denoted Hard Disk Drive (HDD), are composed of magnetic disks that are rotating. Data are read and write thanks to a moving magnetic head. Reads and writes will generate a rotation and a magnetic head movement, which consumes time.
	SSD (Solid-State Drives) are not constructed like that. There is no magnetic head neither magnetic disks. Instead, data is stored in flash memory cells. Data access is quicker on that kind of disk. Note that SSD also costs more than traditional electromagnetic hard drives.

CPU:
	CPU is the initials of Central Processing Unit. 
	The CPU is also denoted processor. The CPU contains :
		#1 The arithmetic and logic unit (ALU)
		#2 The control unit (CU)
	The CPU is the central component of a computer.

Two types of instructions are performed :
	I/O operations : 
		retrieve numbers stored into memory, store numbers into memory from an the input device (the keyboard), load data from memory, and display it to the user.
	An arithmetic operation : 
		add two numbers.


How to speak to the machine?
	#1 Programming languages are formal languages
		They are two types of programming languages :
			#1 Low level
			#2 High level

	#2 Machine language
		Machine languages are composed exclusively of zeros and ones. 
		An instruction written in a machine language is a suite of 0 and 1.
		Each processor (or family of processors) will define a list of instructions called the instruction set.
	    There is an instruction to add to the number, increment by one, decrement by one, copy data from one location in memory to another place...etc.

		It’s possible to write computer programs directly into machine language. However, this is not easy.

	#3 Assembly language
		The assembly language is a low-level programming language. 
		To create a program in assembly language, 
		the developer will write instruction to one or several files. 
		Those files are named source files.
		Here is an example of an instruction written in x86 Linux assembly :
			// assembly code
			mov    eax,1
			int    0x80
		An assembler is used to convert the source files written in an assembly language to object code files.
		The linker will then transform those object code files into an executable file. 
		An executable file contains all the computer’s necessary instructions to launch the program.

	#4 High-level languages 
		There are plenty of high-level languages on the market, like Go.
		Those languages are not closely bound to machine architecture.
		They offer a convenient way to write instructions. 
		For instance, if we want to make a system call to exit our program, we can write in go :
			os.Exit(1)
		
		In this example, we do not have to move a number into a register;
		we use the languages constructs (functions, packages, methods, variables, types ...).
		High-level programs are also written to files.
		Files are named “source files”. 
		Generally, programming languages require adding a specific extension to the filename. 
		For Go programs we will add “.go” At the end of each file that we will write. 
		In PHP the extension is “.php”.

		When source files are written, the program that they define cannot be executed immediately. 
		The source file needs to be compiled by using a compiler. 
		The compiler will transform source files into an executable. 
		The compiler is also a program. 
		Go is part of the compiled language family.

	#4.1 Compiled vs. Interpreted
		Note that some programming languages are interpreted. When the source files have been written, the programmer does not need to compile the sources. With the source files ready, the system can execute the program thanks to an interpreter. Each instruction written into the source file is translated and executed by the interpreter. In some cases, interpreters are storing in a cache a compiled version of the program to boost performance (the source files are not translated each time). PHP, Python, Ruby, Perl are interpreted languages.


*/

package main

