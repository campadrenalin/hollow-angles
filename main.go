package main

import "github.com/codegangsta/martini"
import "github.com/codegangsta/martini-contrib/render"

func root_handler(r render.Render) {
    r.HTML(200, "root", nil)
}

func main() {
    m := martini.Classic()
    m.Use(render.Renderer())

    m.Get("/", root_handler)

    m.Run()
}
