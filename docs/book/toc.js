// Populate the sidebar
//
// This is a script, and not included directly in the page, to control the total size of the book.
// The TOC contains an entry for each page, so if each page includes a copy of the TOC,
// the total size of the page becomes O(n**2).
class MDBookSidebarScrollbox extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.innerHTML = '<ol class="chapter"><li class="chapter-item expanded "><a href="intro.html"><strong aria-hidden="true">1.</strong> Introduction</a></li><li class="chapter-item expanded "><a href="installation.html"><strong aria-hidden="true">2.</strong> Installation</a></li><li class="chapter-item expanded "><a href="cred_quick-setup.html"><strong aria-hidden="true">3.</strong> Setup</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="cred_quick-setup.html"><strong aria-hidden="true">3.1.</strong> Quick Setup</a></li><li class="chapter-item expanded "><a href="cred_init.html"><strong aria-hidden="true">3.2.</strong> Manual Setup</a></li><li class="chapter-item expanded "><a href="cred_migrate.html"><strong aria-hidden="true">3.3.</strong> Migration</a></li></ol></li><li class="chapter-item expanded "><a href="config.html"><strong aria-hidden="true">4.</strong> Configuration</a></li><li class="chapter-item expanded "><a href="cred.html"><strong aria-hidden="true">5.</strong> CLI Reference</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="cred_env.html"><strong aria-hidden="true">5.1.</strong> Environment Variables (cred env)</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="cred_env_copy.html"><strong aria-hidden="true">5.1.1.</strong> Copy</a></li><li class="chapter-item expanded "><a href="cred_env_edit.html"><strong aria-hidden="true">5.1.2.</strong> Edit</a></li><li class="chapter-item expanded "><a href="cred_env_insert.html"><strong aria-hidden="true">5.1.3.</strong> Insert</a></li><li class="chapter-item expanded "><a href="cred_env_ls.html"><strong aria-hidden="true">5.1.4.</strong> List</a></li><li class="chapter-item expanded "><a href="cred_env_mv.html"><strong aria-hidden="true">5.1.5.</strong> Move</a></li><li class="chapter-item expanded "><a href="cred_env_rm.html"><strong aria-hidden="true">5.1.6.</strong> Remove</a></li><li class="chapter-item expanded "><a href="cred_env_show.html"><strong aria-hidden="true">5.1.7.</strong> Show</a></li></ol></li><li class="chapter-item expanded "><a href="cred_pass.html"><strong aria-hidden="true">5.2.</strong> Passwords (cred pass)</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="cred_pass_copy.html"><strong aria-hidden="true">5.2.1.</strong> Copy</a></li><li class="chapter-item expanded "><a href="cred_pass_edit.html"><strong aria-hidden="true">5.2.2.</strong> Edit</a></li><li class="chapter-item expanded "><a href="cred_pass_generate.html"><strong aria-hidden="true">5.2.3.</strong> Generate</a></li><li class="chapter-item expanded "><a href="cred_pass_insert.html"><strong aria-hidden="true">5.2.4.</strong> Insert</a></li><li class="chapter-item expanded "><a href="cred_pass_ls.html"><strong aria-hidden="true">5.2.5.</strong> List</a></li><li class="chapter-item expanded "><a href="cred_pass_mv.html"><strong aria-hidden="true">5.2.6.</strong> Move</a></li><li class="chapter-item expanded "><a href="cred_pass_rm.html"><strong aria-hidden="true">5.2.7.</strong> Remove</a></li><li class="chapter-item expanded "><a href="cred_pass_show.html"><strong aria-hidden="true">5.2.8.</strong> Show</a></li></ol></li></ol></li><li class="chapter-item expanded "><a href="cred_completion.html"><strong aria-hidden="true">6.</strong> Shell Completions</a></li><li><ol class="section"><li class="chapter-item expanded "><a href="cred_completion_bash.html"><strong aria-hidden="true">6.1.</strong> Bash</a></li><li class="chapter-item expanded "><a href="cred_completion_zsh.html"><strong aria-hidden="true">6.2.</strong> Zsh</a></li><li class="chapter-item expanded "><a href="cred_completion_fish.html"><strong aria-hidden="true">6.3.</strong> Fish</a></li><li class="chapter-item expanded "><a href="cred_completion_powershell.html"><strong aria-hidden="true">6.4.</strong> PowerShell</a></li></ol></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString().split("#")[0];
        if (current_page.endsWith("/")) {
            current_page += "index.html";
        }
        var links = Array.prototype.slice.call(this.querySelectorAll("a"));
        var l = links.length;
        for (var i = 0; i < l; ++i) {
            var link = links[i];
            var href = link.getAttribute("href");
            if (href && !href.startsWith("#") && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The "index" page is supposed to alias the first chapter in the book.
            if (link.href === current_page || (i === 0 && path_to_root === "" && current_page.endsWith("/index.html"))) {
                link.classList.add("active");
                var parent = link.parentElement;
                if (parent && parent.classList.contains("chapter-item")) {
                    parent.classList.add("expanded");
                }
                while (parent) {
                    if (parent.tagName === "LI" && parent.previousElementSibling) {
                        if (parent.previousElementSibling.classList.contains("chapter-item")) {
                            parent.previousElementSibling.classList.add("expanded");
                        }
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', function(e) {
            if (e.target.tagName === 'A') {
                sessionStorage.setItem('sidebar-scroll', this.scrollTop);
            }
        }, { passive: true });
        var sidebarScrollTop = sessionStorage.getItem('sidebar-scroll');
        sessionStorage.removeItem('sidebar-scroll');
        if (sidebarScrollTop) {
            // preserve sidebar scroll position when navigating via links within sidebar
            this.scrollTop = sidebarScrollTop;
        } else {
            // scroll sidebar to current active section when navigating via "next/previous chapter" buttons
            var activeSection = document.querySelector('#sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        var sidebarAnchorToggles = document.querySelectorAll('#sidebar a.toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(function (el) {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define("mdbook-sidebar-scrollbox", MDBookSidebarScrollbox);
