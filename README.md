# jhelp

### Purpose
Do you ever wish you could have a quick list of cli commands available in the command line?
jhelp is a simple cli for managing a list of cli tool commands!

This is named 'jhelp' in honor of Joe Franks for the original idea.

### Install
Run
```
go install github.com/timreimherr/jhelp
```
Then
```
jhelp
```

### Adding to the list
New Section
```
jhelp addSection <section-name>
```

New action/command under a specific section
```
jhelp addInfo <sectionId> <info-action> <info-command>
```
