package functions

import "fmt"

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	const dockerfileLines = `RUN apt-get install -y \
  curl \
  libicu57 \
  libunwind8 \
  unzip \
  wget

RUN curl -s https://api.github.com/repos/azure/azure-functions-core-tools/releases/latest \
  | grep "browser_download_url.*linux-x64.*zip\"" \
  | cut -d : -f 2,3 \
  | tr -d \" \
  | wget -qi -

RUN unzip Azure.Functions.*.zip -d /func-cli
RUN chmod +x /func-cli/func

RUN ln -s /func-cli/func /usr/bin/func 
	`

	fmt.Fprintf(m.Out, dockerfileLines)
	return nil
}
