var CircleProgressBar = {
    New: function (canvas, bgColor, color, lineWidth, percent, text) {
        if (text === undefined) {
            text = (percent * 100) + "%";
        }
        var bar = {};
        var ctx = canvas.getContext("2d");
        var radius = canvas.height / 2;
        ctx.translate(radius, radius);
        radius = radius * 0.90;
        ctx.lineWidth = lineWidth;

        bar.drawCircle = function (percent, color) {
            ctx.beginPath();
            var from = 1.5 * Math.PI;
            var to = (percent - 0.25) * Math.PI * 2;
            if (percent<0.25) {
                to = (1.5 + percent * 2) * Math.PI;
            }
            ctx.arc(0, 0, radius, from , to);
            ctx.strokeStyle = color;
            ctx.stroke()
        };

        bar.drawText = function (text) {
            ctx.font="30px Verdana";
            // Create gradient
            var gradient=ctx.createLinearGradient(0,0,ctx.measureText(text).width,0);
            gradient.addColorStop("0","magenta");
            gradient.addColorStop("0.5","blue");
            gradient.addColorStop("1.0","red");
            // Fill with gradient
            ctx.fillStyle=gradient;
            ctx.fillText(text,-ctx.measureText(text).width/2, 15);
        };

        bar.drawCircle(1.0, bgColor); // Draw background circle
        bar.drawCircle(percent, color); // Draw progress bar
        bar.drawText(text);

        return bar;
    }
};