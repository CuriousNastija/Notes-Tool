# Notes Tool

## Description

The **Notes Tool** is a command-line application designed to help users manage notes organized into collections. Users can create, view, add, and remove notes within specific collections.


## Usage

- To use, create or manage a collection named "Your_Note", you would run:

```bash
$ ./notestool Your_Note
```
- After creating a new collection it asks you to set a new password. After you have created a new password you have to restart the program and use the password you created.

## Example

- To use, create or manage a collection named "Your_Note", you would run:

```bash
$ ./notestool Your_Note
```

- When you launch Notes Tool, you will see a menu of operations:

```bash
$ ./Welcome to the notes tool!


Select operation:
1. Show notes.        //Select option `1` to view all notes i n the current collection.
2. Add a note.        //Select option `2` and input the text of the new note when prompted. the note will be added to the collection.
3. Delete a note.     //Choose option `3`, then enter the number of the note you wish to remove. Enter `0`to cancel this operation.
4. Exit.              //To close the application, select option `4`.

```

Here are examples of how an interaction with the tool looks like:

```bash
$ ./notestool
This tool helps manage a collection of notes. You can view, add and delete notes in the specified collection. If the collection doesn't exist, it will be created automatically.

Usage ./notestool <collection_name>

Read README.md to learn more about the tool
```
```bash
$ ./notestool Your_Note
Welcome to the notes tool!

Select operation: (1/4)
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit

Notes:
001 - note hello
002 - note hi
```

```bash
$ ./notestool Your_Note
Welcome to the notes tool!

Select operation: (2/4)
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit

Enter the note text:
note heya

Select operation: (1/4)
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit

Notes:
001 note hello
002 note hi
003 note heya
```

```bash
$ ./notestool Your_note
Welcome to the notes tool!

Select operation: (3/4)
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit

Enter the number of note to remove or 0 to cancel:
3

Select operation: (1/4)
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit

Notes:
001 - note hello
002 - note hi
```

## Data Storage
Each note you make is stored in a plain text file (for example: hello.txt). Notes are saved as a new line within file.

## Allowed Packages
For this project, the following packages are allowed to use:

 - [bufio](https://pkg.go.dev/bufio): Implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O. 

 - [fmt](https://pkg.go.dev/fmt): Implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler. 

 - [os](https://pkg.go.dev/os): Provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information. 

 - [strconv](https://pkg.go.dev/strconv): Implements conversions to and from string representations of basic data types. 

 - [strings](https://pkg.go.dev/strings) : Implements simple functions to manipulate UTF-8 encoded strings. 

 - [time](https://pkg.go.dev/time) : Package time provides functionality for measuring and displaying time.