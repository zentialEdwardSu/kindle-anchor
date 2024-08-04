package main

// Origin Author: chaishushan@gmail.com

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	flagRootDir   = flag.String("d", "", "webdav root dir")
	flagHttpAddr  = flag.String("http", ":8080", "http or https address")
	flagHttpsMode = flag.Bool("https-mode", false, "use https mode")
	flagVerbose   = flag.Bool("Verbose", false, "Show prased params")
	flagCertFile  = flag.String("https-cert-file", "cert.pem", "https cert file")
	flagKeyFile   = flag.String("https-key-file", "key.pem", "https key file")
	flagUserName  = flag.String("user", "", "user name")
	flagPassword  = flag.String("password", "", "user password")
	flagReadonly  = flag.Bool("read-only", false, "read only")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of WebDAV Server\n")
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()
	if *flagVerbose {
		fmt.Println("Server Configuration:")
		fmt.Printf("Root Directory: %s\n", *flagRootDir)
		fmt.Printf("HTTP Address: %s\n", *flagHttpAddr)
		fmt.Printf("HTTPS Mode: %v\n", *flagHttpsMode)
		fmt.Printf("Certificate File: %s\n", *flagCertFile)
		fmt.Printf("Key File: %s\n", *flagKeyFile)
		fmt.Printf("Username: %s\n", *flagUserName)
		fmt.Printf("Password: %s\n", *flagPassword)
		fmt.Printf("Read-Only Mode: %v\n", *flagReadonly)
	}

	fs := &webdav.Handler{
		FileSystem: webdav.Dir(*flagRootDir),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if *flagUserName != "" && *flagPassword != "" {
			username, password, ok := req.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if username != *flagUserName || password != *flagPassword {
				http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
				return
			}
		}
		if req.Method == "GET" && handleDirList(fs.FileSystem, w, req) {
			return
		}
		if *flagReadonly {
			switch req.Method {
			case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
				http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
				return
			}
		}
		fs.ServeHTTP(w, req)
	})
	if *flagHttpsMode {
		err := http.ListenAndServeTLS(*flagHttpAddr, *flagCertFile, *flagKeyFile, nil)
		if err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	} else {
		err := http.ListenAndServe(*flagHttpAddr, nil)
		if err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}

}

func handleDirList(fs webdav.FileSystem, w http.ResponseWriter, req *http.Request) bool {
	ctx := context.Background()
	f, err := fs.OpenFile(ctx, req.URL.Path, os.O_RDONLY, 0)
	if err != nil {
		return false
	}
	defer f.Close()
	if fi, _ := f.Stat(); fi != nil && !fi.IsDir() {
		return false
	}
	dirs, err := f.Readdir(-1)
	if err != nil {
		log.Print(w, "Error reading directory", http.StatusInternalServerError)
		return false
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	for _, d := range dirs {
		name := d.Name()
		if d.IsDir() {
			name += "/"
		}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", name, name)
	}
	fmt.Fprintf(w, "</pre>\n")
	return true
}
