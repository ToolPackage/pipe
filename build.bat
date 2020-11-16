git describe --all --long | pipe run in=text.cut(-8)=regexp.replace('buildId=\w{8}', '')=$id;\
    in.file('version.go')=regexp.replace("buildId=\w{8}", $id)=out.file('version.go')
