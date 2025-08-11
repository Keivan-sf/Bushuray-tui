FROM archlinux:base

# Install Go and build tools
RUN pacman -Sy --noconfirm go git base-devel && pacman -Scc --noconfirm

# Set workdir
WORKDIR /src

# Copy your source code
COPY . .

# Build binary
RUN go build -o bushuray

# Default command: show binary info
CMD ["file", "./bushuray"]
