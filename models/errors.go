package models

import "errors"

var ErrNotFound = errors.New("pokémon introuvable")
var ErrBadRequest = errors.New("numéro invalide")
var ErrTimeout = errors.New("temps dépassé")
var ErrUnavailable = errors.New("service indisponible")
