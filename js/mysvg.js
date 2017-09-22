var TreeNode = {
    New: function (tag) {
        var node = {};
        node.tag = tag;
        node.x = 0;
        node.y = 0;
        node.width = 200;
        node.height = 100;

        node.shape = "rectangle";
        node.text = "";
        // methods
        node.svg = function () {

        };
        return node;
    }
};

var Ellipse = {
    New: function () {
        node = TreeNode.New("ellipse");
        node.xRaidus = 0;
        node.yRaidus = 0;

    }
};