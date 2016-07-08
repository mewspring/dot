gen:
	if [ ! -d "graphviz" ]; then \
		git clone https://github.com/ellson/graphviz.git; \
	fi

	# Copy *.gv and *.dot files.
	find graphviz -type f -name '*.gv' -not -wholename "graphviz/rtest/share/b545.gv" -not -wholename "graphviz/rtest/imagepath_test/base.gv" | xargs -I '{}' cp "{}" .
	find graphviz -type f -name '*.dot' | xargs -I '{}' cp "{}" .

	# Rename *.gv to *.dot.
	rename .gv .dot *.gv

	# Remove execute permissions.
	chmod 0644 *.dot

	# Convert Latin1 encoded files to UTF-8.
	grep -l "charset=latin1" *.dot | xargs -I '{}' recode ISO-8859-1..UTF8 "{}"
	recode ISO-8859-1..UTF8 Latin1.dot

	rm -rf graphviz

.PHONY: gen
