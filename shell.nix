{ pkgs ? import <nixpkgs> {} }:

let
  find_dc = pkgs.buildGoModule rec {
    pname = "find_dc";
    version = "1.0.2";

    src = pkgs.fetchFromGitHub {
      owner = "0xthefr34k";
      repo = "find_dc";
      rev = "v${version}";
      hash = "sha256-rPw0H2jlCyeAKGLCR1Gvisv+hHCCmyRWi1k+i5zzxTg=";
    };

    vendorHash = null;

    meta = {
      description = "find_dc is a Go-based tool designed to process the output of the nxc command, identify domain controllers, and generate the necessary /etc/hosts entries.";
      homepage = "https://github.com/0xthefr34k/find_dc";
      license = pkgs.lib.licenses.mit;
      maintainers = [ "0xthefr34k" ];
    };
  };
in

pkgs.mkShell {
  buildInputs = [ find_dc ];
  shellHook = ''
    echo "find_dc is available!"
  '';
}

