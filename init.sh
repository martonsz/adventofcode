#!/bin/bash
set -e
set -o pipefail

if [[ -d .venv ]]; then
    # shellcheck disable=SC1091
    if [[ -d .venv/bin ]]; then
      source ".venv/bin/activate"
    else
      source ".venv/Scripts/activate"
    fi
else
    python3 -m venv .venv
    # shellcheck disable=SC1091
    if [[ -d .venv/bin ]]; then
      source ".venv/bin/activate"
    else
      source ".venv/Scripts/activate"
    fi
    python -m pip install --upgrade pip
    pip install -r requirements.txt
fi
