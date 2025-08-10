{
  description = "Bushuray";

  inputs = { nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05"; };

  outputs = { self, nixpkgs }:
    let
      allSystems = [
        "x86_64-linux"
      ];
      forAllSystems = f:
        nixpkgs.lib.genAttrs allSystems
        (system: f { pkgs = import nixpkgs { inherit system; }; });
    in {
      packages = forAllSystems ({ pkgs }: {
        # regular nix build (NixOS-only)
        default = pkgs.buildGo124Module rec {
          pname = "bushuray";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-ucw4elcWAEqMa9HOFLqbBlYkgBZ8COn+M0WB1RlymhY=";
        };

        # portable static binary (musl)
        portable = pkgs.pkgsCross.musl64.buildGo124Module rec {
          pname = "bushuray";
          version = "0.1.0";
          src = ./.;
          vendorHash = "sha256-ucw4elcWAEqMa9HOFLqbBlYkgBZ8COn+M0WB1RlymhY=";
        };
      });
    };
}
