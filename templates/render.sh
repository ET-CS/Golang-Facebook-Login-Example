#!/bin/bash

cd js
./render.sh
cd ..
cd sass
./render.sh
cd ..

# SETTINGS
# set to yes to use Cheetah instead of Jinja2 Template Engine
# This example should work for the both of them.
export UseCheetah=false

# compile
if [ "$UseCheetah" = true ] ; then
    cheetah compile _layout --quiet
    cheetah compile index --quiet
fi

# run
python - <<EOF
import os

# Use Cheetah instead of Jinja2
useCheetah = (os.environ['UseCheetah']=="true")

import codecs

if useCheetah:
    print "Using Cheetah"
    # compile cheetah templates
    from index import index
    namespace = { 'appname': '2TE-Bootstrap Example' }
    html = str(index(searchList=[namespace]))
else:
    print "using Jinja2"
    # compile jinja templates
    from jinja2 import Environment, PackageLoader, FileSystemLoader
    import jinja2

    def include_file(name):
        return jinja2.Markup(loader.get_source(env, name)[0])

    loader = FileSystemLoader(os.path.join(os.path.dirname(__file__), ''))
    env = Environment(loader=loader)
    env.globals['include_file'] = include_file
    template = env.get_template('index.html')
    html = template.render(appname='2TE-Bootstrap Example')

# minify using html_minify
from htmlmin.minify import html_minify
html = html_minify(html)

# save minified version
file = codecs.open("../index.min.html", "w", "utf-8")
file.write(html)
file.close()
EOF

if [ "$UseCheetah" = true ] ; then
    echo Cleaning...
    #clean files
    rm *.pyc
    rm _layout.py
    rm index.py
    rm *.bak -f
fi