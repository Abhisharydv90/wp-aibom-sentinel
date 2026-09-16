<?php
/**
 * Plugin Name: WP-AIBOM Sentinel CLI
 * Description: WP-CLI commands for AIBOM Sentinel.
 * Version: 1.0.0
 * Author: Your Name
 */

if ( defined( 'WP_CLI' ) && WP_CLI ) {

    class AIBOM_CLI_Command {

        /**
         * Scan a file for AI-generated code.
         *
         * ## OPTIONS
         *
         * <file>
         * : Path to the file to scan.
         *
         * ## EXAMPLES
         *
         *     wp aibom scan wp-content/plugins/hello.php
         *
         * @when after_wp_load
         */
        public function scan( $args, $assoc_args ) {
            $file = $args[0];
            
            // Resolve to absolute path so the D: drive executable can find it
            $real_path = realpath( $file );

            if ( ! $real_path ) {
                WP_CLI::error( "File not found: $file" );
            }

            // Path to your compiled Go executable
            $exe_path = 'D:\\wp-aibom-sentinel\\aibom.exe';

            // Execute the Go CLI
            $command = escapeshellarg( $exe_path ) . ' scan ' . escapeshellarg( $real_path );
            $output = shell_exec( $command . ' 2>&1' ); // Capture stderr too

            if ( empty( $output ) ) {
                WP_CLI::error( "Failed to run AIBOM scanner. No output." );
            }

            // Parse the JSON output to create a table
            $json_start = strpos( $output, '{' );
            if ( $json_start !== false ) {
                $json_output = substr( $output, $json_start );
                $data = json_decode( $json_output, true );

                if ( $data && isset( $data['error'] ) ) {
                     WP_CLI::error( $data['error'] );
                }

                if ( $data && isset( $data['evidence'] ) ) {
                    WP_CLI::line( "📄 File: " . $data['file'] );
                    WP_CLI::line( "🤖 AI Probability: " . ( $data['ai_probability'] * 100 ) . "%" );
                    WP_CLI::line( "⚠️  Risk Level: " . $data['risk_level'] );
                    WP_CLI::line( "" );

                    $rows = [];
                    foreach ( $data['evidence'] as $e ) {
                        $rows[] = [
                            'Line'    => $e['line'],
                            'Pattern' => $e['pattern'],
                            'Snippet' => substr( $e['snippet'], 0, 50 ) . '...'
                        ];
                    }
                    WP_CLI\Utils\format_items( 'table', $rows, [ 'Line', 'Pattern', 'Snippet' ] );
                } else {
                    // If it's valid JSON but no evidence, just print raw output
                    WP_CLI::line( $output );
                }
            } else {
                WP_CLI::line( $output );
            }
        }
    }

    WP_CLI::add_command( 'aibom', 'AIBOM_CLI_Command' );
}