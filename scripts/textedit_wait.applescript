on run argv
  set thePath to POSIX file (item 1 of argv)
  tell application "TextEdit"
    activate
    set theDoc to open thePath
    repeat while (exists theDoc)
      delay 0.2
    end repeat
  end tell
end run


