CLASSPATH=".:/usr/local/lib/antlr-4.9-complete.jar:$${CLASSPATH}"
antlr4=java -jar /usr/local/lib/antlr-4.9-complete.jar
grun=java org.antlr.v4.gui.TestRig

.PHONY: qasm3
qasm3:
	$(antlr4) -Dlanguage=Go -o $@ ./$@.g4
