package internal

import (
    "io"
    "os"
)

func mergeFiles(inputs []string, output string) error {
    out, err := os.Create(output)
    if err != nil {
        return err
    }
    defer out.Close()

    for _, f := range inputs {
        file, err := os.Open(f)
        if err != nil {
            continue
        }
        io.Copy(out, file)
        file.Close()
    }

    return nil
}
