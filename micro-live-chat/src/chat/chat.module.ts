import {Module} from '@nestjs/common';
import {RoomsService} from './rooms/rooms.service';
import {ConfigService} from "@nestjs/config";
import {ClientsModule, Transport} from "@nestjs/microservices";
import {join} from "path";
import {TypeOrmModule} from "@nestjs/typeorm";
import {ChatMessage} from "./chat-message.model";
import {RabbitMQModule} from "@golevelup/nestjs-rabbitmq";
import {SaveChatMessageService} from "./save-chat-message/save-chat-message.service";

@Module({
    imports: [
        RabbitMQModule.forRootAsync(RabbitMQModule, { // configuracao rabbitmq
            useFactory: () => ({
                uri: `amqp://${process.env.RABBITMQ_USER}:${process.env.RABBITMQ_PASSWORD}@${process.env.RABBITMQ_HOST}:5672`,
                exchanges: [ // cadastro das exchanges
                    {name: 'chat-message', type: 'fanout'}
                ]
            }),
        }),
        TypeOrmModule.forFeature([ChatMessage]),
        new Promise((resolve) => { // Criando cliente RPC
            setTimeout(() => {
                // registrando conexoes RPC
                resolve(ClientsModule.register([
                    {
                        name: 'LIVE_PACKAGE', // nome para usar no Inject do nestjs
                        transport: Transport.GRPC,
                        options: {
                            url: process.env.MICRO_LIVE_GENERATOR_GRPC_URL,
                            package: 'live',
                            protoPath: join(__dirname, '../proto/live.proto')
                        },
                    },
                    {
                        name: 'BOT_PACKAGE',
                        transport: Transport.GRPC,
                        options: {
                            url: process.env.MICRO_LIVE_BOT_GRPC_URL,
                            package: 'education.code.codeedu',
                            protoPath: join(__dirname, '../proto/bot_message.proto')
                        },
                    },
                ]))
            }, 500)
        })
    ],
    providers: [RoomsService, ConfigService, SaveChatMessageService]
})
export class ChatModule {
}
