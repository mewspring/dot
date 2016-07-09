test: a.dot
	# Copy input DOT test files.
	mkdir -p input
	cp *.dot input/

	# Generate output DOT files.
	mkdir -p output
	cp *.dot output/
	dotfmt -i output/*.dot

	# Generate input PNG images for DOT files smaller than 10 kB.
	find input -type f -size -102400c -name '*.dot' -not -name "b15.dot" -not -name "html4.dot" -not -name "multi.dot" -not -name "sides.dot" -not -name "tee.dot" | sort | xargs -I '{}' dot2png "{}"

	# Generate output PNG images for DOT files smaller than 10 kB.
	find output -type f -size -102400c -name '*.dot' -not -name "b15.dot" -not -name "html4.dot" -not -name "multi.dot" -not -name "sides.dot" -not -name "tee.dot" | sort | xargs -I '{}' dot2png "{}"

	# Compare input and output PNG images.
	ls input/*.png | sort | xargs -I '{}' basename "{}" | xargs -I '{}' imgcmp "input/{}" "output/{}"

a.dot: gen

gen: graphviz
	# Copy *.gv and *.dot files.
	find graphviz -type f -name '*.gv' -not -wholename "graphviz/rtest/share/b545.gv" -not -name "base.gv" | xargs -I '{}' cp "{}" .
	find graphviz -type f -name '*.dot' | xargs -I '{}' cp "{}" .

	# Rename *.gv to *.dot.
	rename .gv .dot *.gv

	# Remove execute permissions.
	chmod 0644 *.dot

	# Convert Latin1 encoded files to UTF-8.
	grep -l "charset=latin1" *.dot | xargs -I '{}' recode ISO-8859-1..UTF8 "{}"
	recode ISO-8859-1..UTF8 Latin1.dot

graphviz:
	git clone https://github.com/ellson/graphviz.git

clean:
	rm -rf *.dot graphviz input output

.PHONY: test gen clean
