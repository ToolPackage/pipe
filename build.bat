git describe --all --long | pipe in=text.cut(-8)=regexp.replace("buildId=\w{8}", "{}")=out.file('version.txt')
