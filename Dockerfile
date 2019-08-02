FROM scratch

ADD build/c4-portfolio /c4-portfolio

EXPOSE 8080

ENTRYPOINT ["./c4-portfolio"]