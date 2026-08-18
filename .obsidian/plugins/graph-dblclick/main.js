/*
 * Graph Double-Click Open
 *
 * Obsidian's graph opens a note on a single click. This plugin suppresses that:
 * a single click does nothing, a double-click (or Cmd/Ctrl-click) opens the note.
 *
 * Note: the graph renderer is not part of Obsidian's public API. The plugin patches
 * `view.renderer.onNodeClick` defensively and restores it on unload; if a future
 * Obsidian release renames that hook, the plugin simply does nothing.
 */

const { Plugin, Notice } = require("obsidian");

const DOUBLE_CLICK_MS = 350;
const GRAPH_VIEWS = ["graph", "localgraph"];

module.exports = class GraphDoubleClickPlugin extends Plugin {
  async onload() {
    this.lastClick = { id: null, at: 0 };
    this.patched = new Set(); // { renderer, original }

    this.app.workspace.onLayoutReady(() => this.patchAll());
    this.registerEvent(this.app.workspace.on("layout-change", () => this.patchAll()));
    this.registerEvent(this.app.workspace.on("active-leaf-change", () => this.patchAll()));

    this.addCommand({
      id: "graph-dblclick-status",
      name: "Graph double-click: show status",
      callback: () => {
        new Notice(
          this.patched.size
            ? `Double-click open active on ${this.patched.size} graph view(s).`
            : "No graph view patched. Open a graph, then run this again."
        );
      },
    });
  }

  patchAll() {
    for (const type of GRAPH_VIEWS) {
      for (const leaf of this.app.workspace.getLeavesOfType(type)) {
        this.patchView(leaf.view);
      }
    }
  }

  patchView(view) {
    const renderer = view && view.renderer;
    if (!renderer || typeof renderer.onNodeClick !== "function") return;
    for (const entry of this.patched) {
      if (entry.renderer === renderer) return; // already patched
    }

    const original = renderer.onNodeClick.bind(renderer);
    const plugin = this;

    renderer.onNodeClick = function (evt, id, type) {
      // Modifier click keeps Obsidian's normal behaviour (new tab / new pane).
      if (evt && (evt.metaKey || evt.ctrlKey || evt.shiftKey || evt.altKey)) {
        plugin.lastClick = { id: null, at: 0 };
        return original(evt, id, type);
      }

      const now = Date.now();
      const isSecondClick =
        plugin.lastClick.id === id && now - plugin.lastClick.at < DOUBLE_CLICK_MS;

      if (isSecondClick) {
        plugin.lastClick = { id: null, at: 0 };
        return original(evt, id, type);
      }

      // First click: remember it, open nothing.
      plugin.lastClick = { id, at: now };
    };

    this.patched.add({ renderer, original });
  }

  onunload() {
    for (const { renderer, original } of this.patched) {
      try {
        renderer.onNodeClick = original;
      } catch (_) {
        /* renderer already destroyed */
      }
    }
    this.patched.clear();
  }
};
