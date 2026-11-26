package style

import ansicode "iris.local/draw-test/Ansi/Codes"

var Bold []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PMode(ansicode.RenderingModeBold)}

var Magenta []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PForeground(ansicode.ColorMagenta)}

var Green []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PForeground(ansicode.ColorGreen)}

var Dim []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PMode(ansicode.RenderingModeDim)}

var Red []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PForeground(ansicode.ColorRed)}

var Cyan []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PForeground(ansicode.ColorCyan)}

var Yellow []ansicode.GraphicsParam = []ansicode.GraphicsParam{ansicode.PForeground(ansicode.ColorYellow)}
