var CircleProgressBar = {
    New: function (canvas, bgColor, circleColor, color, lineWidth, percent, text) {
        var bar = {};
        var ctx = canvas.getContext("2d");
        ctx.fillStyle = bgColor;
        ctx.fillRect(0, 0, canvas.width, canvas.height)
        var radius = canvas.height / 2;
        ctx.translate(radius, radius);
        radius = radius * 0.90;
        ctx.lineWidth = lineWidth;

        bar.drawCircle = function (percent, color) {
            ctx.beginPath();
            var from = 1.5 * Math.PI;
            var to = (percent - 0.25) * Math.PI * 2;
            var remain = percent - 0.25;
            if (remain < 0) {
                to = (1.5 + percent * 2) * Math.PI;
            }else{
                to = 2 * Math.PI;
            }
            ctx.arc(0, 0, radius, from , to);
            if (remain > 0) {
                to = remain * 2 * Math.PI;
                ctx.arc(0, 0, radius, 0, to);
            }
            ctx.strokeStyle = color;
            ctx.stroke();
            ctx.closePath();
        };

        bar.drawText = function (text, color) {
            ctx.font="40px Verdana";
            ctx.fillStyle = bgColor;
            ctx.fillRect(-ctx.measureText(oldText).width/2, -20, ctx.measureText(oldText).width, 42);
            ctx.fillStyle = color; // or whatever color the text should be.
            ctx.fillText(text,-ctx.measureText(text).width/2, 20);
        };

        var oldText = "";
        var pct = 0.0;
        bar.drawAnimated = function () {
            bar.drawCircle(1.0, circleColor);
            pct = 0.0;
            requestAnimationFrame(bar.animation);
        };

        bar.animation = function () {
            pct += 0.01;
            color = 'rgb(' + parseInt((1.0-pct) * 255) + ',' + parseInt(pct * 255) + ', 0)';
            bar.drawCircle(pct, color);
            text = parseInt(pct * 100) + " %";
            bar.drawText(text, color);
            oldText = text;
            if (pct < percent) {
                requestAnimationFrame(bar.animation);
            }
        };

        bar.drawCircle(1.0, circleColor); // Draw background circle

        return bar;
    }
};