#include <tarantool/module.h>
#include <msgpuck.h>
#include <uuid/uuid.h>
#include <string.h>
#include <unistd.h>

#define TOKENS_SPACE_ID 512
#define TOKENS_SPACE_PK 0

int get_new_token(box_function_ctx_t *ctx, const char *args_begin, const char *args_end) {
//    uuid_t uuid;
//    uuid_generate(uuid);
//
    const int uuid_str_len = 36;
//    char uuid_str[uuid_str_len + 1]; // ex. "1b4e28ba-2fa1-11d2-883f-0016d3cca427" + "\0"
//    uuid_unparse_lower(uuid, uuid_str);
//    printf("%s\n", uuid_str);

    char uuid_str[] = "1b4e28ba-2fa1-11d2-883f-0016d3cca427";

    char resp_buffer[64] = {0};
    char *cursor = resp_buffer;

    cursor = mp_encode_array(cursor, 2);
    cursor = mp_encode_str(cursor, uuid_str, uuid_str_len);
    cursor = mp_encode_uint(cursor, 228);

    box_tuple_t *token_tuple;
    int err = box_replace(TOKENS_SPACE_ID, resp_buffer, cursor, &token_tuple);
    if (err != 0) {
        say_error("%s", box_error_message(box_error_last()));
        return 1;
    }

    return box_return_tuple(ctx, token_tuple);
}

int revoke_token(box_function_ctx_t *ctx, char *begin, char *end) {
    box_tuple_t *token_tuple;
    int err = box_delete(TOKENS_SPACE_ID, TOKENS_SPACE_PK, begin, end, &token_tuple);
    if (err != 0) {
        say_error("%s", box_error_message(box_error_last()));
        return 1;
    }

    char response[1];
    char *response_end = mp_encode_bool(response, token_tuple != NULL);
    return box_return_mp(ctx, response, response_end);
}

int check_token(box_function_ctx_t *ctx, char *begin, char *end) {
    box_tuple_t *token_tuple;
    int err = box_index_get(TOKENS_SPACE_ID, TOKENS_SPACE_PK, begin, end, &token_tuple);
    if (err != 0) {
        say_error("%s", box_error_message(box_error_last()));
        return 1;
    }

    char response[1];
    char *response_end = mp_encode_bool(response, token_tuple != NULL);
    return box_return_mp(ctx, response, response_end);
}
