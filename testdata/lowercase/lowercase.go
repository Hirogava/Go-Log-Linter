package lowercase

import "log/slog"

func test() {
    slog.Info("Starting server") // want "must start with lowercase"
}
