require 'middleman-core/renderers/redcarpet'

class CustomMarkdownRenderer < Middleman::Renderers::MiddlemanRedcarpetHTML
  def block_code(code, language)
    if language.nil? || language.empty?
      "<pre><code>#{ERB::Util.html_escape(code)}</code></pre>"
    else
      lexer = Rouge::Lexer.find_fancy(language, code) || Rouge::Lexers::PlainText
      formatter = Rouge::Formatters::HTML.new
      highlighted = formatter.format(lexer.lex(code))
      "<div class=\"highlight #{language} tab-#{language}\"><pre class=\"highlight #{language} tab-#{language}\"><code>#{highlighted}</code></pre></div>"
    end
  end

  def header(text, header_level)
    id = text.downcase.gsub(/<[^>]*>/, '').gsub(/[^\w]+/, '-').gsub(/^-|-$/, '')
    "<h#{header_level} id='#{id}'>#{text}</h#{header_level}>\n"
  end
end

set :source, '.'
set :build_dir, '../docs-html'

set :markdown_engine, :redcarpet
set :markdown,
    renderer: CustomMarkdownRenderer,
    fenced_code_blocks: true,
    smartypants: true,
    disable_indented_code_blocks: true,
    prettify: true,
    strikethrough: true,
    tables: true,
    with_toc_data: true,
    no_intra_emphasis: true

ignore 'Gemfile'
ignore 'Gemfile.lock'
ignore 'config.rb'
ignore /^vendor\/.*/

set :css_dir, 'stylesheets'
set :js_dir, 'javascripts'
set :images_dir, 'images'
set :fonts_dir, 'fonts'

configure :build do
  activate :minify_css
  activate :minify_javascript, compressor: proc { require 'uglifier'; Uglifier.new(harmony: true) }
  activate :asset_hash
end

activate :syntax

# Middleman 4.x dropped Sprockets support. We manually resolve //= require directives.
# asset_hash fingerprints based on file content at build time, so we must write the
# real concatenated content to the source JS files BEFORE the build runs, then restore
# the originals (with //= require stubs) AFTER the build.

JS_SRC = File.expand_path('javascripts', __dir__)

# Strip //= require lines then read file content
read_js = ->(f) { File.read(File.join(JS_SRC, "#{f}.js")).gsub(/^\/\/=.*\n/, '') }

NOSEARCH_FILES = %w[
  lib/_jquery
  lib/_imagesloaded.min
  lib/_energize
  app/_copy
  app/_toc
  app/_lang
].freeze

SEARCH_FILES = %w[
  lib/_jquery
  lib/_imagesloaded.min
  lib/_lunr
  lib/_jquery.highlight
  lib/_energize
  app/_copy
  app/_toc
  app/_lang
].freeze

before_build do
  nosearch_stub = File.read(File.join(JS_SRC, 'all_nosearch.js'))
  search_stub   = File.read(File.join(JS_SRC, 'all.js'))

  # Save originals so after_build can restore them
  @_js_originals = {
    'all_nosearch.js' => nosearch_stub,
    'all.js'          => search_stub
  }

  nosearch_body = nosearch_stub.gsub(/^\/\/=.*\n/, '')
  search_body   = File.read(File.join(JS_SRC, 'app/_search.js')).gsub(/^\/\/=.*\n/, '')

  File.write(File.join(JS_SRC, 'all_nosearch.js'),
    (NOSEARCH_FILES.map(&read_js) + [nosearch_body]).join("\n"))

  File.write(File.join(JS_SRC, 'all.js'),
    (SEARCH_FILES.map(&read_js) + [nosearch_body, search_body]).join("\n"))
end

after_build do
  # Restore //= require stubs so git doesn't see source file changes
  if @_js_originals
    @_js_originals.each do |filename, content|
      File.write(File.join(JS_SRC, filename), content)
    end
  end

  # asset_hash generates root-relative paths (/stylesheets/...) but the binary
  # serves docs at /docs/, so convert to relative paths so assets resolve correctly.
  index_out = File.join('../docs-html', 'index.html')
  if File.exist?(index_out)
    html = File.read(index_out)
    html.gsub!(%r{(href|src)="/(stylesheets|javascripts|fonts|images)/}, '\1="\2/')
    File.write(index_out, html)
  end
end

helpers do
  def toc_data(page_content)
    require 'nokogiri'
    doc = Nokogiri::HTML::DocumentFragment.parse(page_content)
    h1s = []
    doc.css('h1, h2').each do |header|
      if header.name == 'h1'
        h1s << {
          id: header[:id],
          title: header.inner_text,
          content: header.inner_html,
          children: []
        }
      elsif header.name == 'h2' && !h1s.empty?
        h1s.last[:children] << {
          id: header[:id],
          title: header.inner_text,
          content: header.inner_html
        }
      end
    end
    h1s
  end
end
