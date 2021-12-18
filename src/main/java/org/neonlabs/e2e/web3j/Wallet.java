package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import java.math.BigInteger;
import java.security.InvalidAlgorithmParameterException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.util.Random;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import org.neonlabs.e2e.web3j.model.Account;
import org.web3j.crypto.CipherException;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.Keys;
import org.web3j.crypto.WalletFile;

/**
 * Creates a new wallet.
 */
@Slf4j
@SuppressWarnings({ "PMD.GuardLogStatement", "PMD.SystemPrintln" })
public class Wallet {
  /**
   * Creates an externally owned account.
   */
  @SneakyThrows
  @Step
  public static Account create() {

    final var seed = generateSeed();
    var account = new Account();

    try {
      System.out.println("00001");
      final ECKeyPair ecKeyPair = Keys.createEcKeyPair();
      System.out.println("00002");
      final BigInteger privateKeyInDec = ecKeyPair.getPrivateKey();
      System.out.println("00003");

      final String privatekeyInHex = privateKeyInDec.toString(16);
      System.out.println("00004");

      final WalletFile walletFile = org.web3j.crypto.Wallet.createLight(seed, ecKeyPair);
      System.out.println("00005");
      String address = walletFile.getAddress();
      System.out.println("00006");

      account.setAddress(address);
      account.setPrivateKey(privatekeyInHex);
      account.setPrivateKeyDec(privateKeyInDec);
      System.out.println("00007");

    } catch (CipherException e) {
      log.error(e.getClass().getName());
      log.error(e.getMessage());
    } catch (InvalidAlgorithmParameterException e) {
      log.error(e.getClass().getName());
      log.error(e.getMessage());
    } catch (NoSuchAlgorithmException e) {
      log.error(e.getClass().getName());
      log.error(e.getMessage());
    } catch (NoSuchProviderException e) {
      log.error(e.getClass().getName());
      log.error(e.getMessage());
    }

    return account;
  }

  private static String generateSeed() {
    final var seed = IntStream.rangeClosed(1, 12)
        .mapToObj(index -> generateRandomString() + " ")
        .collect(Collectors.joining(" "));
    log.info("Generated seed: %s", seed);
    return seed.trim();
  }

  private static String generateRandomString() {
    final var startCharacter = 97;
    final var endCharacter = 122;
    final var maxLength = 10;
    final var random = new Random();
    return random.ints(startCharacter, endCharacter + 1)
        .limit(maxLength)
        .collect(StringBuilder::new, StringBuilder::appendCodePoint, StringBuilder::append)
        .toString();
  }
}
