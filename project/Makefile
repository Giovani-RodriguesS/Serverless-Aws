.PHONY: build deploy validate # devem ser interpretados como comandos

v:
	sam validate
b:
	sam build
d:
	sam deploy --guided
sy:
	sam sync --stack-name serverless-it --watch
bp:
	sam build && sam deploy